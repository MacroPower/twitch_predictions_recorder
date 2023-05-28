package event

import (
	"time"

	"github.com/MacroPower/twitch_predictions_recorder/internal/eventraw"
)

type EventMixin struct {
	ChannelName string `json:"channel_name"`
}

func ConvertMessage(m *eventraw.Message, cm EventMixin) Event {
	var outcomes []Outcome
	for _, o := range m.Data.Event.Outcomes {
		outcomes = append(outcomes, convertOutcome(m.Data.Event.ID, o))
	}

	return Event{
		ID:                      m.Data.Event.ID,
		ChannelID:               m.Data.Event.ChannelID,
		ChannelName:             cm.ChannelName,
		CreatedAt:               m.Data.Event.CreatedAt,
		CreatedBy:               convertUser(m.Data.Event.CreatedBy),
		EndedAt:                 m.Data.Event.EndedAt,
		EndedBy:                 convertUser(m.Data.Event.EndedBy),
		LockedAt:                m.Data.Event.LockedAt,
		LockedBy:                convertUser(m.Data.Event.LockedBy),
		PredictionWindowSeconds: m.Data.Event.PredictionWindowSeconds,
		Title:                   m.Data.Event.Title,
		EventStates:             []EventState{convertEventState(m)},
		Outcomes:                outcomes,
	}
}

func convertEventState(m *eventraw.Message) EventState {
	var outcomeStates []OutcomeState
	for _, o := range m.Data.Event.Outcomes {
		outcomeStates = append(outcomeStates, convertOutcomeState(o, m.Data.Timestamp))
	}

	return EventState{
		Type:          m.Type,
		Timestamp:     m.Data.Timestamp,
		Status:        m.Data.Event.Status,
		OutcomeStates: outcomeStates,
	}
}

func convertUser(u eventraw.User) User {
	return User{
		UserID:          u.UserID,
		UserDisplayName: u.UserDisplayName,
	}
}

func convertOutcome(event string, o eventraw.Outcome) Outcome {
	return Outcome{
		ID:           o.ID,
		EventID:      event,
		Color:        o.Color,
		Title:        o.Title,
		BadgeVersion: o.Badge.Version,
		BadgeSetID:   o.Badge.SetID,
	}
}

func convertOutcomeState(o eventraw.Outcome, timestamp time.Time) OutcomeState {
	var topPredictors []Predictor
	var resultType string
	for _, p := range o.TopPredictors {
		topPredictors = append(topPredictors, convertPredictor(p))
		resultType = p.Result.Type
	}

	return OutcomeState{
		OutcomeID:     o.ID,
		Timestamp:     timestamp,
		TotalPoints:   o.TotalPoints,
		TotalUsers:    o.TotalUsers,
		TopPredictors: topPredictors,
		ResultType:    resultType,
	}
}

func convertPredictor(p eventraw.Predictor) Predictor {
	return Predictor{
		PredictorID: p.ID,
		EventID:     p.EventID,
		OutcomeID:   p.OutcomeID,
		ChannelID:   p.ChannelID,
		User: User{
			UserID:          p.UserID,
			UserDisplayName: p.UserDisplayName,
		},
		Points:               p.Points,
		PredictedAt:          p.PredictedAt,
		UpdatedAt:            p.UpdatedAt,
		ResultType:           p.Result.Type,
		ResultPointsWon:      p.Result.PointsWon,
		ResultIsAcknowledged: p.Result.IsAcknowledged,
	}
}
