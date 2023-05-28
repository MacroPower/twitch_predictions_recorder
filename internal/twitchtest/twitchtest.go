package twitchtest

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"

	"github.com/MacroPower/twitch_predictions_recorder/internal/event"
	"github.com/MacroPower/twitch_predictions_recorder/internal/eventraw"
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
		msg := eventraw.Message{}
		if err := json.Unmarshal(scanner.Bytes(), &msg); err != nil {
			return fmt.Errorf("failed to unmarshal test data: %w", err)
		}

		if err := dataFunc(event.ConvertMessage(&msg, event.EventMixin{ChannelName: "test"})); err != nil {
			return fmt.Errorf("failed to handle test data: %w", err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to scan test data: %w", err)
	}

	return nil
}
