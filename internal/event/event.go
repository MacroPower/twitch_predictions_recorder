package event

import "time"

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

func (e *EventData) ToSamples(channel string) Samples {
	blueOutcome := e.getOutcomeWithColor("BLUE")
	pinkOutcome := e.getOutcomeWithColor("PINK")

	return Samples{
		Timestamp: e.Timestamp,

		ID:          e.Event.ID,
		ChannelID:   e.Event.ChannelID,
		ChannelName: channel,

		CreatedAt: e.Event.CreatedAt,
		EndedAt:   e.Event.EndedAt,
		LockedAt:  e.Event.LockedAt,

		CreatedBy: e.Event.CreatedBy.UserDisplayName,
		EndedBy:   e.Event.EndedBy.UserDisplayName,
		LockedBy:  e.Event.LockedBy.UserDisplayName,

		BlueID:          blueOutcome.ID,
		BlueTitle:       blueOutcome.Title,
		BlueTotalPoints: blueOutcome.TotalPoints,
		BlueTotalUsers:  blueOutcome.TotalUsers,
		BlueWon:         e.Event.WinningOutcomeID == blueOutcome.ID,

		//BluePredictors:  blueOutcome.TopPredictors,

		PinkID:          pinkOutcome.ID,
		PinkTitle:       pinkOutcome.Title,
		PinkTotalPoints: pinkOutcome.TotalPoints,
		PinkTotalUsers:  pinkOutcome.TotalUsers,
		PinkWon:         e.Event.WinningOutcomeID == pinkOutcome.ID,

		//PinkPredictors:  pinkOutcome.TopPredictors,

		PredictionWindowSeconds: e.Event.PredictionWindowSeconds,
		Status:                  e.Event.Status,
		Title:                   e.Event.Title,
	}
}

type Samples struct {
	Timestamp time.Time

	ID          string
	ChannelID   string
	ChannelName string

	CreatedAt time.Time
	EndedAt   time.Time
	LockedAt  time.Time

	CreatedBy string
	EndedBy   string
	LockedBy  string

	// Outcomes

	BlueID          string
	BlueTitle       string
	BlueTotalPoints int
	BlueTotalUsers  int
	BlueWon         bool
	//BluePredictors  []Predictor

	PinkID          string
	PinkTitle       string
	PinkTotalPoints int
	PinkTotalUsers  int
	PinkWon         bool
	//PinkPredictors  []Predictor

	PredictionWindowSeconds int
	Status                  string // ACTIVE or LOCKED or RESOLVE_PENDING or RESOLVED
	Title                   string
}

func (e *EventData) getOutcomeWithColor(color string) Outcome {
	for _, o := range e.Event.Outcomes {
		if o.Color == color {
			return o
		}
	}

	return Outcome{}
}
