package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDBSettings struct {
	Host    string
	Port    int
	SSLMode string

	User     string
	Password string

	DBName   string
	TimeZone string
}

func NewPostgresDB(pgdb PostgresDBSettings) (*GormDB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d sslmode=%s user=%s password=%s dbname=%s TimeZone=%s",
		pgdb.Host, pgdb.Port, pgdb.SSLMode, pgdb.User, pgdb.Password, pgdb.DBName, pgdb.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening db connection: %w", err)
	}

	return &GormDB{db}, nil
}
