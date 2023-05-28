package api

import (
	"net/http"

	"github.com/MacroPower/twitch_predictions_recorder/internal/log"
	"github.com/MacroPower/twitch_predictions_recorder/ui"
)

type IndexHTTP struct {
	logger log.Logger
}

func NewIndexHTTP(logger log.Logger) *IndexHTTP {
	return &IndexHTTP{logger: logger}
}

func (h *IndexHTTP) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		rawFile, err := ui.StaticFiles.ReadFile("dist/favicon.ico")
		if err != nil {
			log.Error(h.logger).Log("msg", "Error reading file", "err", err)
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
		if _, err := w.Write(rawFile); err != nil {
			log.Error(h.logger).Log("msg", "Error writing response", "err", err)
		}

		return
	}

	rawFile, err := ui.StaticFiles.ReadFile("dist/index.html")
	if err != nil {
		log.Error(h.logger).Log("msg", "Error reading file", "err", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
	if _, err := w.Write(rawFile); err != nil {
		log.Error(h.logger).Log("msg", "Error writing response", "err", err)
	}
}
