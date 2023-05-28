package api

import (
	"fmt"
	"io/fs"
	"net/http"

	"github.com/MacroPower/twitch_predictions_recorder/internal/db"
	"github.com/MacroPower/twitch_predictions_recorder/internal/log"
	"github.com/MacroPower/twitch_predictions_recorder/ui"
)

func Router(db db.DB, logger log.Logger) (*http.ServeMux, error) {
	mux := http.NewServeMux()

	// index page
	mux.Handle("/", NewIndexHTTP(logger))

	// static files
	staticFS, err := fs.Sub(ui.StaticFiles, "dist")
	if err != nil {
		return nil, fmt.Errorf("failed to create static file subtree: %w", err)
	}
	httpFS := http.FileServer(http.FS(staticFS))
	mux.Handle("/static/", httpFS)

	// api
	mux.Handle("/api/v1/summary", NewSummaryHTTP(db, log.With(logger, "endpoint", "summary")))
	mux.Handle("/api/v1/details", NewDetailsHTTP(db, log.With(logger, "endpoint", "details")))

	return mux, nil
}
