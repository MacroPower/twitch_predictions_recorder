package twitch

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/MacroPower/twitch_predictions_recorder/internal/event"
	"github.com/MacroPower/twitch_predictions_recorder/internal/eventraw"
	"github.com/MacroPower/twitch_predictions_recorder/internal/log"

	"github.com/Adeithe/go-twitch"
	"github.com/Adeithe/go-twitch/api"
	"github.com/Adeithe/go-twitch/api/helix"
	"github.com/Adeithe/go-twitch/pubsub"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	appName         = "twitch_predictions_recorder"
	streamerSegSize = 25
	maxShards       = 10
	topic           = "predictions-channel-v1"
)

var (
	messagesProcessed = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: appName,
		Subsystem: "messages",
		Name:      "processed_total",
		Help:      "The total number of messages processed",
	}, []string{
		"streamer",
		"shard",
		"status",
	})

	shardConnected = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: appName,
		Subsystem: "shard",
		Name:      "connected",
		Help:      "1 if the shard is connected, 0 otherwise",
	}, []string{
		"shard",
	})

	shardUpdates = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: appName,
		Subsystem: "shard",
		Name:      "updates_total",
		Help:      "The total number of shard updates",
	}, []string{
		"shard",
	})

	shardReconnections = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: appName,
		Subsystem: "shard",
		Name:      "reconnections_total",
		Help:      "The total number of shard reconnections",
	}, []string{
		"shard",
	})

	shardLatency = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: appName,
		Subsystem: "shard",
		Name:      "latency_seconds",
		Help:      "The shard latency in seconds",
	}, []string{
		"shard",
	})
)

type EventListener struct {
	twitchClient *pubsub.Client
	apiClient    *api.Client
	streamers    []string
	logger       log.Logger
}

func NewEventListener(client *api.Client, logger log.Logger, streamers ...string) *EventListener {
	mgr := twitch.PubSub()
	mgr.SetMaxShards(maxShards)

	for i := 0; i < maxShards; i++ {
		for _, j := range streamers {
			for _, k := range []string{"ACTIVE", "LOCKED", "RESOLVE_PENDING", "RESOLVED"} {
				messagesProcessed.WithLabelValues(j, fmt.Sprint(i), k)
			}
		}
		shardUpdates.WithLabelValues(fmt.Sprint(i))
		shardReconnections.WithLabelValues(fmt.Sprint(i))
		shardLatency.WithLabelValues(fmt.Sprint(i))
		shardConnected.WithLabelValues(fmt.Sprint(i))
	}

	mgr.OnShardConnect(func(shard int) {
		shardConnected.WithLabelValues(fmt.Sprint(shard)).Set(1)
		log.Info(logger).Log("msg", "Shard connected", "shard", shard)
	})

	mgr.OnShardReconnect(func(shard int) {
		shardReconnections.WithLabelValues(fmt.Sprint(shard)).Inc()
		log.Info(logger).Log("msg", "Shard reconnected", "shard", shard)
	})

	mgr.OnShardLatencyUpdate(func(shard int, latency time.Duration) {
		shardUpdates.WithLabelValues(fmt.Sprint(shard)).Inc()
		shardLatency.WithLabelValues(fmt.Sprint(shard)).Set(latency.Seconds())
		log.Info(logger).Log("msg", "Shard updated", "shard", shard, "ping_ms", latency.Milliseconds())
	})

	mgr.OnShardDisconnect(func(shard int) {
		shardConnected.WithLabelValues(fmt.Sprint(shard)).Set(0)
		log.Info(logger).Log("msg", "Shard disconnected", "shard", shard)
	})

	return &EventListener{
		twitchClient: mgr,
		apiClient:    client,
		streamers:    streamers,
		logger:       logger,
	}
}

func (te *EventListener) Listen(dataFunc func(event.Event) error) error {
	streamerIDName := te.GetIDMap(te.streamers...)

	te.twitchClient.OnShardMessage(func(shard int, topic string, data []byte) {
		msg := &eventraw.Message{}
		if err := json.Unmarshal(data, msg); err != nil {
			log.Error(te.logger).Log("msg", "Error unmarshalling event", "err", err.Error())

			return
		}

		streamer := streamerIDName[msg.Data.Event.ChannelID]

		messagesProcessed.WithLabelValues(streamer, fmt.Sprint(shard), msg.Data.Event.Status).Inc()
		log.Debug(te.logger).Log(
			"msg", "Got message",
			"streamer", streamer,
			"shard", shard,
			"topic", topic,
			"type", msg.Type,
			"status", msg.Data.Event.Status,
		)
		if err := dataFunc(event.ConvertMessage(msg, event.EventMixin{ChannelName: streamer})); err != nil {
			log.Error(te.logger).Log("err", err)
		}
	})

	for id := range streamerIDName {
		if err := te.twitchClient.Listen(topic, id); err != nil {
			return fmt.Errorf("failed to listen on topic: %w", err)
		}
	}

	log.Info(te.logger).Log(
		"msg", "Started listening",
		"topics", te.twitchClient.GetNumTopics(),
		"shards", te.twitchClient.GetNumShards(),
	)

	return nil
}

func (te *EventListener) GetIDMap(streamers ...string) map[string]string {
	streamersSeg := make([][]string, (len(streamers)/streamerSegSize)+1)
	for i := 0; i < len(streamers); i++ {
		streamersSeg[i/streamerSegSize] = append(streamersSeg[i/streamerSegSize], streamers[i])
	}

	ids := make(map[string]string)

	for i, seg := range streamersSeg {
		start, end := getSize(i, streamerSegSize, len(streamers))
		log.Info(te.logger).Log("msg", fmt.Sprintf("Getting IDs for streamer batch %d (%d - %d)", i+1, start, end))

		ud, err := te.apiClient.Helix().GetUsers(helix.UserOpts{Logins: seg})
		if err != nil {
			panic(err)
		}

		for _, d := range ud.Data {
			ids[d.ID] = d.DisplayName
			log.Debug(te.logger).Log("msg", "Got ID for streamer", "name", d.DisplayName, "id", d.ID)
		}
	}

	return ids
}

func (te *EventListener) Close() {
	te.twitchClient.Close()
}

func getSize(iter int, size int, max int) (int, int) {
	start := iter * size
	end := (iter + 1) * size

	if end > max {
		return start, max
	}

	return start, end - 1
}
