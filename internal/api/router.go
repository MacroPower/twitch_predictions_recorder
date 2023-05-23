package api

import (
	"encoding/json"
	"io/fs"
	"net/http"

	"github.com/MacroPower/twitch_predictions_recorder/internal/db"
	"github.com/MacroPower/twitch_predictions_recorder/internal/log"
	"github.com/MacroPower/twitch_predictions_recorder/ui"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		rawFile, _ := ui.StaticFiles.ReadFile("dist/favicon.ico")
		w.Write(rawFile)
		return
	}
	rawFile, _ := ui.StaticFiles.ReadFile("dist/index.html")
	w.Write(rawFile)
}

type SummaryHTTP struct {
	db     db.DB
	logger log.Logger
}

func NewSummaryHTTP(db db.DB, logger log.Logger) *SummaryHTTP {
	return &SummaryHTTP{db: db, logger: logger}
}

func (s *SummaryHTTP) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query()
	summary, _, err := s.db.GetSummary(query.Get("id"))
	if err != nil {
		log.Error(s.logger).Log("msg", "Error getting summary", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(summary)
	if err != nil {
		log.Error(s.logger).Log("msg", "Error marshaling summary", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

type DetailsHTTP struct {
	db     db.DB
	logger log.Logger
}

func NewDetailsHTTP(db db.DB, logger log.Logger) *DetailsHTTP {
	return &DetailsHTTP{db: db, logger: logger}
}

func (h *DetailsHTTP) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query()
	id := query.Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "id is required"}`))
		return
	}
	details, _, err := h.db.GetDetails(id)
	if err != nil {
		log.Error(h.logger).Log("msg", "Error getting details", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(details)
	if err != nil {
		log.Error(h.logger).Log("msg", "Error marshaling details", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func Router(db db.DB, logger log.Logger) (*http.ServeMux, error) {
	mux := http.NewServeMux()

	// index page
	mux.HandleFunc("/", indexHandler)

	// static files
	staticFS, err := fs.Sub(ui.StaticFiles, "dist")
	if err != nil {
		return nil, err
	}
	httpFS := http.FileServer(http.FS(staticFS))
	mux.Handle("/static/", httpFS)

	// api
	mux.Handle("/api/v1/summary", NewSummaryHTTP(db, logger))
	mux.Handle("/api/v1/details", NewDetailsHTTP(db, logger))

	return mux, nil
}
