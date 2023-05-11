package twitchtest

import (
	"bufio"
	"encoding/json"
	"io"

	"github.com/MacroPower/twitch_predictions_recorder/internal/event"
)

type TestListener struct {
	Reader io.Reader
}

func NewTestListener(r io.Reader) *TestListener {
	return &TestListener{
		Reader: r,
	}
}

func (l *TestListener) Listen(dataFunc func(event.Event) error) error {
	scanner := bufio.NewScanner(l.Reader)

	for scanner.Scan() {
		msg := event.Message{}
		if err := json.Unmarshal(scanner.Bytes(), &msg); err != nil {
			return err
		}

		if err := dataFunc(event.Event{
			StreamerName: "test",
			Message:      &msg,
		}); err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
