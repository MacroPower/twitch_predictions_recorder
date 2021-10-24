package db

import (
	"fmt"

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

func (pgdb PostgresDB) NewDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d sslmode=%s user=%s password=%s dbname=%s TimeZone=%s",
		pgdb.Host, pgdb.Port, pgdb.SSLMode, pgdb.User, pgdb.Password, pgdb.DBName, pgdb.TimeZone,
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func Setup(gdb *gorm.DB, obj interface{}) {
	if err := gdb.AutoMigrate(obj); err != nil {
		panic(err)
	}
}

func SetupHypertable(gdb *gorm.DB) {
	statement := "create_hypertable('samples', 'timestamp')"
	gdb.Select(statement)
}

func SetupDefault(gdb *gorm.DB) {
	Setup(gdb, &Samples{})
	SetupHypertable(gdb)
	Setup(gdb, &Predictor{})
}
