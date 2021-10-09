package db

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	Host    string
	Port    int
	SSLMode string

	User     string
	Password string

	DBName   string
	TimeZone string
}

type Samples struct {
	Time  time.Time
	Value float64
}

func (db PostgresDB) NewDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d sslmode=%s user=%s password=%s dbname=%s TimeZone=%s",
		db.Host, db.Port, db.SSLMode, db.User, db.Password, db.DBName, db.TimeZone,
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func Setup(db *gorm.DB, name string, obj interface{}) {
	if err := db.AutoMigrate(obj); err != nil {
		panic(err)
	}
	db.Select("create_hypertable('samples', 'time')")
}
