package db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSqliteDB(path string) (*GormDB, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening db connection: %w", err)
	}

	return &GormDB{db}, nil
}
