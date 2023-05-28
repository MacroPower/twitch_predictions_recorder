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

var testDB *db.GormDB

func init() {
	var err error
	testDB, err = db.NewSqliteDB(dbPath)
	if err != nil {
		panic(err)
	}
}

func TestSqlite(t *testing.T) {
	t.Parallel()

	err := testDB.SetupDefaults()
	require.NoError(t, err)

	file, err := os.Open("testdata/prediction-1.json")
	require.NoError(t, err)

	lt := twitchtest.NewTestListener(bufio.NewReader(file))
	err = lt.Listen(func(e event.Event) error {
		err = testDB.AddEvents(e)
		require.NoError(t, err)

		return nil
	})
	require.NoError(t, err)
}

//nolint:paralleltest
func TestGet(t *testing.T) {
	_, _, err := testDB.GetSummary("")
	require.NoError(t, err)
}

//nolint:paralleltest
func TestGetDetails(t *testing.T) {
	_, _, err := testDB.GetDetails("08773786-e48a-4df2-8759-db007c3f7a64")
	require.NoError(t, err)
}

//nolint:paralleltest
func TestGetSimilar(t *testing.T) {
	got, _, err := testDB.GetRelatedSummaries("Will moon Survive the Third Region?", 3)
	require.NoError(t, err)

	require.Len(t, got, 2)
}
