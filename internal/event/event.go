package event

import (
	"time"
)

type Event struct {
	ID                      string       `json:"id" gorm:"primarykey"`
	ChannelID               string       `json:"channel_id"`
	ChannelName             string       `json:"channel_name"`
	CreatedAt               time.Time    `json:"created_at"`
	CreatedByID             string       `json:"created_by_id"`
	CreatedBy               User         `json:"created_by" gorm:"foreignKey:CreatedByID"`
	EndedAt                 time.Time    `json:"ended_at"`
	EndedByID               string       `json:"ended_by_id"`
	EndedBy                 User         `json:"ended_by" gorm:"foreignKey:EndedByID"`
	LockedAt                time.Time    `json:"locked_at"`
	LockedByID              string       `json:"locked_by_id"`
	LockedBy                User         `json:"locked_by" gorm:"foreignKey:LockedByID"`
	PredictionWindowSeconds int          `json:"prediction_window_seconds"`
	Title                   string       `json:"title"`
	EventStates             []EventState `json:"event_state" gorm:"foreignKey:EventID"`
}

type EventState struct {
	ID               uint      `json:"id" gorm:"primarykey;autoIncrement"`
	EventID          string    `json:"event_id"`
	Timestamp        time.Time `json:"timestamp"`
	Type             string    `json:"type"`
	Outcomes         []Outcome `json:"outcomes" gorm:"foreignKey:EventStateID"`
	Status           string    `json:"status"` // ACTIVE or LOCKED or RESOLVE_PENDING or RESOLVED
	WinningOutcomeID string    `json:"winning_outcome_id"`
}

type User struct {
	UserID          string `json:"user_id" gorm:"primarykey"`
	UserDisplayName string `json:"user_display_name"`
}

type Outcome struct {
	ID            string         `json:"id" gorm:"primarykey"`
	EventStateID  uint           `json:"event_state_id"`
	Color         string         `json:"color"`
	Title         string         `json:"title"`
	BadgeVersion  string         `json:"badge_version"`
	BadgeSetID    string         `json:"badge_set_id"`
	OutcomeStates []OutcomeState `json:"outcome_state" gorm:"foreignKey:OutcomeID"`
}

type OutcomeState struct {
	ID            uint        `json:"id" gorm:"primarykey;autoIncrement"`
	OutcomeID     string      `json:"outcome_id"`
	Timestamp     time.Time   `json:"timestamp"`
	TotalPoints   int         `json:"total_points"`
	TotalUsers    int         `json:"total_users"`
	TopPredictors []Predictor `json:"top_predictors" gorm:"foreignKey:OutcomeStateID"`
}

type Predictor struct {
	ID              string           `json:"id" gorm:"primarykey"`
	EventID         string           `json:"event_id"`
	OutcomeID       string           `json:"outcome_id"`
	OutcomeStateID  uint             `json:"outcome_state_id"`
	ChannelID       string           `json:"channel_id"`
	PredictedByID   string           `json:"predicted_by_id"`
	User            User             `json:"user" gorm:"foreignKey:PredictedByID"`
	PredictorStates []PredictorState `json:"predictor_state" gorm:"foreignKey:PredictorID"`
}

type PredictorState struct {
	ID                   uint      `json:"id" gorm:"primarykey;autoIncrement"`
	PredictorID          string    `json:"predictor_id"`
	Points               int       `json:"points"`
	PredictedAt          time.Time `json:"predicted_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	ResultType           string    `json:"result_type"` // WIN or LOSE
	ResultPointsWon      int       `json:"result_points_won"`
	ResultIsAcknowledged bool      `json:"result_is_acknowledged"`
}
