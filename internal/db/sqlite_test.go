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

	err = db.SetupDefaults()
	require.NoError(t, err)

	file, err := os.Open("testdata/prediction-1.json")
	require.NoError(t, err)

	lt := twitchtest.NewTestListener(bufio.NewReader(file))
	err = lt.Listen(func(e event.Event) error {
		err = db.AddEvents(e)
		require.NoError(t, err)

		return nil
	})
	require.NoError(t, err)
}

//nolint:paralleltest
func TestGet(t *testing.T) {
	db, err := db.NewSqliteDB(dbPath)
	require.NoError(t, err)

	_, _, err = db.GetSummary("")
	require.NoError(t, err)
}

//nolint:paralleltest
func TestGetDetails(t *testing.T) {
	db, err := db.NewSqliteDB(dbPath)
	require.NoError(t, err)

	_, _, err = db.GetDetails("08773786-e48a-4df2-8759-db007c3f7a64")
	require.NoError(t, err)
}
