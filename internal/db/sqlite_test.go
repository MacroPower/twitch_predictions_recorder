package db_test

import (
	"bufio"
	"os"
	"testing"

	"github.com/MacroPower/twitch_predictions_recorder/internal/db"
	"github.com/MacroPower/twitch_predictions_recorder/internal/event"
	"github.com/MacroPower/twitch_predictions_recorder/internal/twitchtest"
	"github.com/stretchr/testify/require"
)

const dbPath = "test.sqlite"

func TestSqlite(t *testing.T) {
	t.Parallel()

	db, err := db.NewSqliteDB(dbPath)
	require.NoError(t, err)

	db.SetupDefaults()

	file, err := os.Open("testdata/prediction-1.json")
	require.NoError(t, err)

	lt := twitchtest.NewTestListener(bufio.NewReader(file))
	lt.Listen(func(e event.Event) error {
		db.AddEvents(e)

		return nil
	})
}

func TestGet(t *testing.T) {
	db, err := db.NewSqliteDB(dbPath)
	require.NoError(t, err)

	_, _, err = db.GetSummary()
	require.NoError(t, err)
}
