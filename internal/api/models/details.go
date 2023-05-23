package models

import "time"

type EventDetails struct {
	ID                      string        `json:"id"`
	ChannelName             string        `json:"channel_name"`
	CreatedAt               time.Time     `json:"created_at"`
	PredictionWindowSeconds int           `json:"prediction_window_seconds"`
	Title                   string        `json:"title"`
	EventSeries             []EventSeries `json:"event_series"`
}

type EventSeries struct {
	ID        string           `json:"id"`
	Timestamp time.Time        `json:"timestamp"`
	Status    string           `json:"status"`
	Outcomes  []OutcomeDetails `json:"outcomes"`
}

type OutcomeDetails struct {
	ID            string      `json:"id"`
	Color         string      `json:"color"`
	Title         string      `json:"title"`
	BadgeVersion  string      `json:"badge_version"`
	BadgeSetID    string      `json:"badge_set_id"`
	Timestamp     time.Time   `json:"timestamp"`
	TotalPoints   int         `json:"total_points"`
	TotalUsers    int         `json:"total_users"`
	ResultType    string      `json:"result_type"`
	TopPredictors []Predictor `json:"top_predictors"`
}

type Predictor struct {
	User   User `json:"user"`
	Points int  `json:"points"`
}
