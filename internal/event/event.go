package event

import (
	"time"
)

type Event struct {
	Type string    `json:"type"`
	Data EventData `json:"data"`
}

type User struct {
	Type              string      `json:"type"`
	UserID            string      `json:"user_id"`
	UserDisplayName   string      `json:"user_display_name"`
	ExtensionClientID interface{} `json:"extension_client_id"`
}

type Predictor struct {
	ID          string    `json:"id"`
	EventID     string    `json:"event_id"`
	OutcomeID   string    `json:"outcome_id"`
	ChannelID   string    `json:"channel_id"`
	Points      int       `json:"points"`
	PredictedAt time.Time `json:"predicted_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserID      string    `json:"user_id"`
	Result      struct {
		Type           string `json:"type"` // WIN or LOSE
		PointsWon      int    `json:"points_won"`
		IsAcknowledged bool   `json:"is_acknowledged"`
	} `json:"result"`
	UserDisplayName string `json:"user_display_name"`
}

type Outcome struct {
	ID            string      `json:"id"`
	Color         string      `json:"color"`
	Title         string      `json:"title"`
	TotalPoints   int         `json:"total_points"`
	TotalUsers    int         `json:"total_users"`
	TopPredictors []Predictor `json:"top_predictors"`
	Badge         struct {
		Version string `json:"version"`
		SetID   string `json:"set_id"`
	} `json:"badge"`
}

type EventData struct {
	Timestamp time.Time `json:"timestamp"`
	Event     struct {
		ID                      string    `json:"id"`
		ChannelID               string    `json:"channel_id"`
		CreatedAt               time.Time `json:"created_at"`
		CreatedBy               User      `json:"created_by"`
		EndedAt                 time.Time `json:"ended_at"`
		EndedBy                 User      `json:"ended_by"`
		LockedAt                time.Time `json:"locked_at"`
		LockedBy                User      `json:"locked_by"`
		Outcomes                []Outcome `json:"outcomes"`
		PredictionWindowSeconds int       `json:"prediction_window_seconds"`
		Status                  string    `json:"status"` // ACTIVE or LOCKED or RESOLVE_PENDING or RESOLVED
		Title                   string    `json:"title"`
		WinningOutcomeID        string    `json:"winning_outcome_id"`
	} `json:"event"`
}
