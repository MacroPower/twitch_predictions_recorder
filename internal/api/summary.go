package api

import (
	"encoding/json"
	"net/http"

	"github.com/MacroPower/twitch_predictions_recorder/internal/api/models"
	"github.com/MacroPower/twitch_predictions_recorder/internal/db"
	"github.com/MacroPower/twitch_predictions_recorder/internal/log"
)

const editDistance = 3

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
	queryID := query.Get("id")
	queryTitle := query.Get("title")

	var summary []models.EventSummary
	var err error
	if queryTitle != "" {
		summary, _, err = s.db.GetRelatedSummaries(queryTitle, editDistance)
	} else {
		summary, _, err = s.db.GetSummary(queryID)
	}
	if err != nil {
		log.Error(s.logger).Log("msg", "Error getting data", "err", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	data, err := json.Marshal(summary)
	if err != nil {
		log.Error(s.logger).Log("msg", "Error marshaling data", "err", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
	if _, err = w.Write(data); err != nil {
		log.Error(s.logger).Log("msg", "Error writing response", "err", err)
	}
}
