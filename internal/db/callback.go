package db

import "github.com/MacroPower/twitch_predictions_recorder/internal/event"

type CallbackDB struct {
	Callback func(...event.Event)
}

func (cdb *CallbackDB) SetupDefaults() {
}

func (cdb *CallbackDB) AddEvents(events ...event.Event) {
	cdb.Callback(events...)
}
