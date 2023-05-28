package db

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSqliteDB(path string) (*GormDB, error) {
	//nolint:dogsled
	_, filename, _, _ := runtime.Caller(0)

	extPath := filepath.Join(filepath.Dir(filename), "lib/distlib/distlib_64")

	sql.Register("sqlite3_advanced",
		&sqlite3.SQLiteDriver{
			Extensions: []string{extPath},
		})

	f := sqlite.Dialector{
		DriverName: "sqlite3_advanced",
		DSN:        path,
	}

	db, err := gorm.Open(f, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening db connection: %w", err)
	}

	return &GormDB{db}, nil
}
