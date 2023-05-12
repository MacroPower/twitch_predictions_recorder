package db

import (
	"github.com/MacroPower/twitch_predictions_recorder/internal/event"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GormDB struct {
	db *gorm.DB
}

type DB interface {
	SetupDefaults()
	AddEvents(...event.Event)
}

func (gdb *GormDB) SetupDefaults() {
	gdb.setup(&event.Event{})
	gdb.setup(&event.EventState{})
	gdb.setup(&event.Outcome{})
	gdb.setup(&event.Predictor{})
	gdb.setup(&event.User{})
}

func (gdb *GormDB) AddEvents(events ...event.Event) {
	gdb.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Table("events").Create(events)
}

func (gdb *GormDB) setup(obj interface{}) {
	if err := gdb.db.AutoMigrate(obj); err != nil {
		panic(err)
	}
}
