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

func NewPostgresDB(pgdb PostgresDB) (*GormDB, error) {
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

type DB interface {
	SetupDefaults()
	AddSamples(samples ...Samples)
}

type GormDB struct {
	db *gorm.DB
}

func (gdb *GormDB) SetupDefaults() {
	gdb.setup(&Samples{})
	gdb.setup(&Predictor{})

	statement := "create_hypertable('samples', 'timestamp')"
	gdb.db.Select(statement)
}

func (gdb *GormDB) AddSamples(samples ...Samples) {
	gdb.db.Table("samples").Create(samples)
}

func (gdb *GormDB) setup(obj interface{}) {
	if err := gdb.db.AutoMigrate(obj); err != nil {
		panic(err)
	}
}

type TestDB struct {
	TestFunc func(...Samples)
}

func (tdb *TestDB) SetupDefaults() {
}

func (tdb *TestDB) AddSamples(samples ...Samples) {
	tdb.TestFunc(samples...)
}
