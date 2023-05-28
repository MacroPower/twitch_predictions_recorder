package api

import (
	"encoding/json"
	"net/http"

	"github.com/MacroPower/twitch_predictions_recorder/internal/db"
	"github.com/MacroPower/twitch_predictions_recorder/internal/log"
)

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
		if _, err := w.Write([]byte(`{"error": "id is required"}`)); err != nil {
			log.Error(h.logger).Log("msg", "Error writing response", "err", err)
		}

		return
	}

	details, _, err := h.db.GetDetails(id)
	if err != nil {
		log.Error(h.logger).Log("msg", "Error getting data", "err", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	data, err := json.Marshal(details)
	if err != nil {
		log.Error(h.logger).Log("msg", "Error marshaling data", "err", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
	if _, err = w.Write(data); err != nil {
		log.Error(h.logger).Log("msg", "Error writing response", "err", err)
	}
}
