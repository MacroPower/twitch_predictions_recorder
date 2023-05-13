package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/MacroPower/twitch_predictions_recorder/internal/db"
	"github.com/MacroPower/twitch_predictions_recorder/internal/event"
	"github.com/MacroPower/twitch_predictions_recorder/internal/log"
	"github.com/MacroPower/twitch_predictions_recorder/internal/twitch"

	"github.com/alecthomas/kong"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	appName = "twitch_predictions_recorder"
)

var cli struct {
	Twitch struct {
		ClientID string `help:"Twitch Client ID." required:""`
		Secret   string `help:"Twitch Secret." required:""`

		Streamers     []string `help:"List of streamers to monitor."`
		StreamersFile string   `help:"List of streamers to monitor." default:"streamers.txt"`
	} `prefix:"twitch." embed:""`

	Database struct {
		Type string `help:"Database type. One of: [postgres, sqlite, test]" default:"sqlite"`

		Postgres struct {
			Host     string `help:"PG Host." default:"info"`
			Port     int    `help:"PG Port." default:"5432"`
			SSLMode  string `help:"PG SSL Mode." default:"prefer"`
			User     string `help:"PG User."`
			Password string `help:"PG Password."`
			DBName   string `help:"PG DB Name." default:"postgres"`
		} `prefix:"pg." embed:""`

		SQLite struct {
			Path string `help:"Path to SQLite database." default:"twitch_predictions_recorder.sqlite"`
		} `prefix:"sqlite." embed:""`

		TimeZone string `help:"Time zone name." default:"America/New_York"`
	} `prefix:"db." embed:""`

	Metrics struct {
		Disable bool          `help:"Disable metrics."`
		Path    string        `help:"Path to serve metrics on." default:"/metrics"`
		Address string        `help:"Address to serve metrics on." default:":8080"`
		Timeout time.Duration `help:"HTTP timeout." default:"60s"`
	} `prefix:"metrics." embed:""`

	Log struct {
		Level  string `help:"Log level." default:"info"`
		Format string `help:"Log format. One of: [logfmt, json]" default:"logfmt"`
	} `prefix:"log." embed:""`
}

func main() {
	cliCtx := kong.Parse(
		&cli,
		kong.Name(appName),
		kong.DefaultEnvars(""),
		kong.Configuration(kong.JSON, ".env.json"),
	)

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

	var err error
	streamers := cli.Twitch.Streamers
	if len(streamers) == 0 {
		log.Info(logger).Log("msg", "No streamers provided via arguments, reading from file")
		streamers, err = getStreamersFromFile(cli.Twitch.StreamersFile, logger)
		if err != nil {
			panic(err)
		}
	}
	if len(streamers) == 0 {
		log.Error(logger).Log("msg", "No streamers to monitor, stopping")

		os.Exit(1)
	}

	api, err := twitch.NewAPIClient(cli.Twitch.ClientID, cli.Twitch.Secret)
	if err != nil {
		log.Error(logger).Log("msg", "Failed to create Twitch API client", "err", err)

		os.Exit(1)
	}

	var gdb db.DB
	switch cli.Database.Type {
	case "postgres":
		gdb, err = db.NewPostgresDB(db.PostgresDBSettings{
			Host:     cli.Database.Postgres.Host,
			Port:     cli.Database.Postgres.Port,
			SSLMode:  cli.Database.Postgres.SSLMode,
			User:     cli.Database.Postgres.User,
			Password: cli.Database.Postgres.Password,
			DBName:   cli.Database.Postgres.DBName,
			TimeZone: cli.Database.TimeZone,
		})
		if err != nil {
			panic(err)
		}
	case "sqlite":
		gdb, err = db.NewSqliteDB(cli.Database.SQLite.Path)
		if err != nil {
			panic(err)
		}
	case "test":
		gdb = &db.CallbackDB{
			Callback: func(events ...event.Event) {
				jsonData, err := json.Marshal(events)
				if err != nil {
					log.Error(logger).Log("msg", "Could not marshal sample data", "err", err)
				}

				log.Debug(logger).Log("msg", "Got events", "data", string(jsonData))
			},
		}
	}

	gdb.SetupDefaults()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)

	listener := twitch.NewEventListener(api, logger, streamers...)
	err = listener.Listen(func(d event.Event) error {
		gdb.AddEvents(d)
		return nil
	})
	if err != nil {
		panic(err)
	}

	if !cli.Metrics.Disable {
		http.Handle(cli.Metrics.Path, promhttp.Handler())
		go func() {
			log.Info(logger).Log("msg", "Starting metrics handler")
			s := &http.Server{
				Addr:         cli.Metrics.Address,
				ReadTimeout:  cli.Metrics.Timeout,
				WriteTimeout: cli.Metrics.Timeout,
			}
			if err := s.ListenAndServe(); err != nil {
				log.Error(logger).Log("msg", "Error serving metrics", "err", err)
			}
			log.Info(logger).Log("msg", "Metrics handler terminated")
		}()
	}

	<-sc
	log.Info(logger).Log("msg", "Stopping")
	listener.Close()
}

func getStreamersFromFile(file string, logger log.Logger) ([]string, error) {
	fileBytes, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	splitFunc := func(c rune) bool {
		return c == '\n'
	}

	streamers := strings.FieldsFunc(strings.ReplaceAll(string(fileBytes), "\r", ""), splitFunc)

	return streamers, nil
}
