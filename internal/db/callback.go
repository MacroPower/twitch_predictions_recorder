package db

import (
	"github.com/MacroPower/twitch_predictions_recorder/internal/api/models"
	"github.com/MacroPower/twitch_predictions_recorder/internal/event"
)

type CallbackDB struct {
	Callback func(...event.Event) error
}

func (cdb *CallbackDB) SetupDefaults() error {
	return nil
}

func (cdb *CallbackDB) AddEvents(events ...event.Event) error {
	return cdb.Callback(events...)
}

func (cdb *CallbackDB) GetSummary() ([]models.EventSummary, string, error) {
	return nil, "", nil
}
