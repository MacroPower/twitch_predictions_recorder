package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/Adeithe/go-twitch"
	"github.com/Adeithe/go-twitch/api/helix"
	"golang.org/x/oauth2/clientcredentials"
	oauth2 "golang.org/x/oauth2/twitch"

	"github.com/MacroPower/twitch_predictions_recorder/internal/db"
	"github.com/MacroPower/twitch_predictions_recorder/internal/event"
)

var (
	clientID = flag.String("client-id", "", "Client ID")
	secret   = flag.String("secret", "", "Secret")

	pgHost      = flag.String("pg-host", "127.0.0.1", "")
	pgPort      = flag.Int("pg-port", 5432, "")
	pgSSLMode   = flag.String("pg-sslmode", "disable", "")
	pgUser      = flag.String("pg-user", "postgres", "")
	pgPassword  = flag.String("pg-password", "", "")
	pgDBName    = flag.String("pg-dbname", "timescale", "")
	pgTimeZone  = flag.String("pg-timezone", "America/New_York", "")
	pgTableName = flag.String("pg-tablename", "samples", "")

	streamersFile = flag.String("streamers-file", "streamers.txt", "List of streamers to monitor")
)

const (
	streamerSegSize = 25
)

func main() {
	flag.Parse()

	fileBytes, err := ioutil.ReadFile(*streamersFile)
	if err != nil {
		panic(err)
	}

	splitFunc := func(c rune) bool {
		return c == '\n'
	}

	streamers := strings.FieldsFunc(strings.ReplaceAll(string(fileBytes), "\r", ""), splitFunc)
	streamersSeg := make([][]string, (len(streamers)/streamerSegSize)+1)
	for i := range streamers {
		streamersSeg[i/streamerSegSize] = append(streamersSeg[i/streamerSegSize], streamers[i])
	}

	api := twitch.API(*clientID)

	oauth2Config := &clientcredentials.Config{
		ClientID:     *clientID,
		ClientSecret: *secret,
		TokenURL:     oauth2.Endpoint.TokenURL,
	}

	token, err := oauth2Config.Token(context.Background())
	if err != nil {
		panic(err)
	}

	api = api.NewBearer(token.AccessToken)

	ids := make(map[string]string)

	for i, seg := range streamersSeg {
		start, end := getSize(i, streamerSegSize, len(streamers))
		fmt.Printf("Getting IDs for streamer batch %d (%d - %d)\n", i+1, start, end)

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

	pgdb, err := db.PostgresDB{
		Host:     *pgHost,
		Port:     *pgPort,
		SSLMode:  *pgSSLMode,
		User:     *pgUser,
		Password: *pgPassword,
		DBName:   *pgDBName,
		TimeZone: *pgTimeZone,
	}.NewDB()
	if err != nil {
		panic(err)
	}

	db.SetupDefault(pgdb)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)

	mgr := twitch.PubSub()

	mgr.OnShardConnect(func(shard int) {
		fmt.Printf("Shard #%d connected!\n", shard)
	})

	mgr.OnShardReconnect(func(shard int) {
		fmt.Printf("Shard #%d reconnected!\n", shard)
	})

	mgr.OnShardMessage(func(shard int, topic string, data []byte) {
		e := &event.Event{}
		if err := json.Unmarshal(data, e); err != nil {
			fmt.Println(err.Error())
		}

		fmt.Printf("%s: %s\n", e.Type, e.Data.Event.ID)
		if e.Data.Event.Status == "RESOLVED" {
			fmt.Printf("Shard #%d > %s %+v\n", shard, topic, e)

			pgdb.Table(*pgTableName).Create([]db.Samples{
				db.ToSamples(&e.Data, ids[e.Data.Event.ChannelID]),
			})
		}
	})

	mgr.OnShardLatencyUpdate(func(shard int, latency time.Duration) {
		fmt.Printf("Shard #%d has %.3fs ping!\n", shard, latency.Seconds())
	})

	mgr.OnShardDisconnect(func(shard int) {
		fmt.Printf("Shard #%d disconnected!\n", shard)
	})

	for id := range ids {
		printErr(mgr.Listen("predictions-channel-v1", id))
	}

	fmt.Printf("Started listening to %d topics on %d shards!\n", mgr.GetNumTopics(), mgr.GetNumShards())

	<-sc
	fmt.Println("Stopping...")
	mgr.Close()
}

func printErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func getSize(iter int, size int, max int) (int, int) {
	start := iter * size
	end := (iter + 1) * size

	if end > max {
		return start, max
	}

	return start, end - 1
}
