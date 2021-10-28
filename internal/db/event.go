package db

import (
	"time"

	"github.com/MacroPower/twitch_predictions_recorder/internal/event"
)

type Samples struct {
	Timestamp time.Time

	EventID     string `gorm:"primarykey"`
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

	PinkID          string
	PinkTitle       string
	PinkTotalPoints int
	PinkTotalUsers  int
	PinkWon         bool

	Predictors []Predictor `gorm:"foreignKey:EventID"`

	PredictionWindowSeconds int
	Status                  string // ACTIVE or LOCKED or RESOLVE_PENDING or RESOLVED
	Title                   string
}

type Outcome struct {
	OutcomeID     string
	Color         string
	Title         string
	TotalPoints   int
	TotalUsers    int
	TopPredictors []Predictor
	BadgeVersion  string
	BadgeSetID    string
}

type Predictor struct {
	PredictorID          string `gorm:"primarykey"`
	EventID              string
	OutcomeID            string
	ChannelID            string
	Color                string
	Points               int
	PredictedAt          time.Time
	UpdatedAt            time.Time
	UserID               string
	ResultType           string // WIN or LOSE
	ResultPointsWon      int
	ResultIsAcknowledged bool
	UserDisplayName      string
}

func ToSamples(e *event.Data, channel string) Samples {
	blueOutcome := getOutcomeWithColor(e, "BLUE")
	pinkOutcome := getOutcomeWithColor(e, "PINK")

	return Samples{
		Timestamp: e.Timestamp,

		EventID:     e.Event.ID,
		ChannelID:   e.Event.ChannelID,
		ChannelName: channel,

		CreatedAt: e.Event.CreatedAt,
		EndedAt:   e.Event.EndedAt,
		LockedAt:  e.Event.LockedAt,

		CreatedBy: e.Event.CreatedBy.UserDisplayName,
		EndedBy:   e.Event.EndedBy.UserDisplayName,
		LockedBy:  e.Event.LockedBy.UserDisplayName,

		BlueID:          blueOutcome.OutcomeID,
		BlueTitle:       blueOutcome.Title,
		BlueTotalPoints: blueOutcome.TotalPoints,
		BlueTotalUsers:  blueOutcome.TotalUsers,
		BlueWon:         e.Event.WinningOutcomeID == blueOutcome.OutcomeID,

		PinkID:          pinkOutcome.OutcomeID,
		PinkTitle:       pinkOutcome.Title,
		PinkTotalPoints: pinkOutcome.TotalPoints,
		PinkTotalUsers:  pinkOutcome.TotalUsers,
		PinkWon:         e.Event.WinningOutcomeID == pinkOutcome.OutcomeID,

		Predictors: mergePredictors(blueOutcome.TopPredictors, pinkOutcome.TopPredictors),

		PredictionWindowSeconds: e.Event.PredictionWindowSeconds,
		Status:                  e.Event.Status,
		Title:                   e.Event.Title,
	}
}

func getOutcomeWithColor(e *event.Data, color string) Outcome {
	for _, o := range e.Event.Outcomes {
		if o.Color == color {
			return Outcome{
				OutcomeID:     o.ID,
				Color:         o.Color,
				Title:         o.Title,
				TotalPoints:   o.TotalPoints,
				TotalUsers:    o.TotalUsers,
				TopPredictors: getPredictors(o.TopPredictors, o.Color),
				BadgeVersion:  o.Badge.Version,
				BadgeSetID:    o.Badge.SetID,
			}
		}
	}

	return Outcome{}
}

func getPredictors(ep []event.Predictor, color string) (result []Predictor) {
	for _, p := range ep {
		result = append(result, Predictor{
			PredictorID:          p.ID,
			EventID:              p.EventID,
			OutcomeID:            p.OutcomeID,
			ChannelID:            p.ChannelID,
			Color:                color,
			Points:               p.Points,
			PredictedAt:          p.PredictedAt,
			UpdatedAt:            p.UpdatedAt,
			UserID:               p.UserID,
			ResultType:           p.Result.Type,
			ResultPointsWon:      p.Result.PointsWon,
			ResultIsAcknowledged: p.Result.IsAcknowledged,
			UserDisplayName:      p.UserDisplayName,
		})
	}

	return result
}

func mergePredictors(eps ...[]Predictor) (result []Predictor) {
	for _, ep := range eps {
		result = append(result, ep...)
	}

	return result
}
