package models

import "time"

type EventSummary struct {
	ID                      string           `json:"id"`
	Timestamp               time.Time        `json:"timestamp"`
	ChannelName             string           `json:"channel_name"`
	PredictionWindowSeconds int              `json:"prediction_window_seconds"`
	Title                   string           `json:"title"`
	Status                  string           `json:"status"`
	Outcomes                []OutcomeSummary `json:"outcomes"`
}

type OutcomeSummary struct {
	Color        string `json:"color"`
	Title        string `json:"title"`
	BadgeVersion string `json:"badge_version"`
	BadgeSetID   string `json:"badge_set_id"`
	TotalPoints  int    `json:"total_points"`
	TotalUsers   int    `json:"total_users"`
	ResultType   string `json:"result_type"`
}
