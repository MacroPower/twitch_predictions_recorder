package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/MacroPower/twitch_predictions_recorder/internal/db"
	"github.com/MacroPower/twitch_predictions_recorder/internal/event"
	"github.com/MacroPower/twitch_predictions_recorder/internal/log"

	twitch "github.com/Adeithe/go-twitch"
	"github.com/Adeithe/go-twitch/api/helix"
	"github.com/alecthomas/kong"
	"golang.org/x/oauth2/clientcredentials"
	oauth2 "golang.org/x/oauth2/twitch"
)

const (
	appName          = "twitch_predictions_recorder"
	streamersSegSize = 25
)

var cli struct {
	TwitchClientID string `help:"Twitch Client ID."`
	TwitchSecret   string `help:"Twitch Secret."`

	StreamersFile string `help:"List of streamers to monitor." default:"streamers.txt"`

	Database struct {
		Type string `help:"Database type. One of: [postgres, test]" default:"postgres"`

		Postgres struct {
			Host     string `help:"PG Host." default:"info"`
			Port     int    `help:"PG Port." default:"5432"`
			SSLMode  string `help:"PG SSL Mode." default:"prefer"`
			User     string `help:"PG User."`
			Password string `help:"PG Password."`
			DBName   string `help:"PG DB Name." default:"postgres"`
		} `prefix:"pg." embed:""`

		TimeZone string `help:"Time zone name." default:"America/New_York"`
	} `prefix:"db." embed:""`

	Log struct {
		Level  string `help:"Log level." default:"info"`
		Format string `help:"Log format. One of: [logfmt, json]" default:"logfmt"`
	} `prefix:"log." embed:""`
}

func main() {
	cliCtx := kong.Parse(&cli, kong.Name(appName), kong.DefaultEnvars(""))

	logLevel := &log.AllowedLevel{}
	if err := logLevel.Set(cli.Log.Level); err != nil {
		cliCtx.FatalIfErrorf(err)
	}

	logFormat := &log.AllowedFormat{}
	if err := logFormat.Set(cli.Log.Format); err != nil {
		cliCtx.FatalIfErrorf(err)
	}

	logger := log.New(&log.Config{
		Level:  logLevel,
		Format: logFormat,
	})

	log.Info(logger).Log("msg", fmt.Sprintf("Starting %s", appName))

	fileBytes, err := os.ReadFile(cli.StreamersFile)
	if err != nil {
		panic(err)
	}

	splitFunc := func(c rune) bool {
		return c == '\n'
	}

	streamers := strings.FieldsFunc(strings.ReplaceAll(string(fileBytes), "\r", ""), splitFunc)
	streamersSeg := make([][]string, (len(streamers)/streamersSegSize)+1)
	for i := range streamers {
		streamersSeg[i/streamersSegSize] = append(streamersSeg[i/streamersSegSize], streamers[i])
	}

	api := twitch.API(cli.TwitchClientID)

	oauth2Config := &clientcredentials.Config{
		ClientID:     cli.TwitchClientID,
		ClientSecret: cli.TwitchSecret,
		TokenURL:     oauth2.Endpoint.TokenURL,
	}

	token, err := oauth2Config.Token(context.Background())
	if err != nil {
		panic(err)
	}

	api = api.NewBearer(token.AccessToken)

	ids := make(map[string]string)

	for i, seg := range streamersSeg {
		start, end := getSize(i, streamersSegSize, len(streamers))
		log.Info(logger).Log("msg", fmt.Sprintf("Getting IDs for streamer batch %d (%d - %d)", i+1, start, end))

		ud, err := api.Helix().GetUsers(helix.UserOpts{
			Logins: seg,
		})
		if err != nil {
			panic(err)
		}

		for _, d := range ud.Data {
			ids[d.ID] = d.DisplayName
		}
	}

	var gdb db.DB

	switch cli.Database.Type {
	case "postgres":
		pgdb := db.PostgresDB{
			Host:     cli.Database.Postgres.Host,
			Port:     cli.Database.Postgres.Port,
			SSLMode:  cli.Database.Postgres.SSLMode,
			User:     cli.Database.Postgres.User,
			Password: cli.Database.Postgres.Password,
			DBName:   cli.Database.Postgres.DBName,
			TimeZone: cli.Database.TimeZone,
		}
		gdb, err = db.NewPostgresDB(pgdb)
		if err != nil {
			panic(err)
		}

	case "test":
		gdb = &db.TestDB{
			TestFunc: func(samples ...db.Samples) {
				jsonData, err := json.Marshal(samples)
				if err != nil {
					log.Error(logger).Log("msg", "Could not marshal sample data", "err", err)
				}

				log.Debug(logger).Log("msg", "Got samples", "data", string(jsonData))
			},
		}
	}

	gdb.SetupDefaults()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)

	mgr := twitch.PubSub()

	mgr.OnShardConnect(func(shard int) {
		log.Info(logger).Log("msg", "Shard connected", "shard", shard)
	})

	mgr.OnShardReconnect(func(shard int) {
		log.Info(logger).Log("msg", "Shard reconnected", "shard", shard)
	})

	mgr.OnShardMessage(func(shard int, topic string, data []byte) {
		msg := &event.Message{}
		if err := json.Unmarshal(data, msg); err != nil {
			log.Error(logger).Log("msg", "Error unmarshalling event", "err", err.Error())

			return
		}

		e := msg.Data.Event
		channel := ids[e.ChannelID]

		log.Debug(logger).Log(
			"msg", "Got message",
			"channel", channel,
			"shard", shard,
			"topic", topic,
			"type", msg.Type,
			"status", e.Status,
		)

		if e.Status == "RESOLVED" {
			gdb.AddSamples(db.ToSamples(&msg.Data, channel))
		}
	})

	mgr.OnShardLatencyUpdate(func(shard int, latency time.Duration) {
		log.Info(logger).Log("msg", "Shard updated", "shard", shard, "ping_ms", latency.Milliseconds())
	})

	mgr.OnShardDisconnect(func(shard int) {
		log.Info(logger).Log("msg", "Shard disconnected", "shard", shard)
	})

	for id := range ids {
		err := mgr.Listen("predictions-channel-v1", id)
		if err != nil {
			log.Error(logger).Log("err", err)
		}
	}

	log.Info(logger).Log("msg", "Started listening", "topics", mgr.GetNumTopics(), "shards", mgr.GetNumShards())

	<-sc
	log.Info(logger).Log("msg", "Stopping")
	mgr.Close()
}

func getSize(iter int, size int, max int) (int, int) {
	start := iter * size
	end := (iter + 1) * size

	if end > max {
		return start, max
	}

	return start, end - 1
}
