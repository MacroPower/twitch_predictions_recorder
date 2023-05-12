package event

import "github.com/MacroPower/twitch_predictions_recorder/internal/eventraw"

type EventMixin struct {
	ChannelName string `json:"channel_name"`
}

func ConvertMessage(m *eventraw.Message, cm EventMixin) Event {
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
	}
}

func convertEventState(m *eventraw.Message) EventState {
	var outcomes []Outcome
	for _, o := range m.Data.Event.Outcomes {
		outcomes = append(outcomes, convertOutcome(o, m.Data.Event.ID))
	}

	return EventState{
		Type:             m.Type,
		Timestamp:        m.Data.Timestamp,
		Outcomes:         outcomes,
		Status:           m.Data.Event.Status,
		WinningOutcomeID: m.Data.Event.WinningOutcomeID,
	}
}

func convertUser(u eventraw.User) User {
	return User{
		UserID:          u.UserID,
		UserDisplayName: u.UserDisplayName,
	}
}

func convertOutcome(o eventraw.Outcome, eventID string) Outcome {
	var topPredictors []Predictor
	for _, p := range o.TopPredictors {
		topPredictors = append(topPredictors, convertPredictor(p))
	}

	return Outcome{
		ID:            o.ID,
		EventID:       eventID,
		Color:         o.Color,
		Title:         o.Title,
		TotalPoints:   o.TotalPoints,
		TotalUsers:    o.TotalUsers,
		TopPredictors: topPredictors,
		BadgeVersion:  o.Badge.Version,
		BadgeSetID:    o.Badge.SetID,
	}
}

func convertPredictor(p eventraw.Predictor) Predictor {
	return Predictor{
		ID:          p.ID,
		EventID:     p.EventID,
		OutcomeID:   p.OutcomeID,
		ChannelID:   p.ChannelID,
		Points:      p.Points,
		PredictedAt: p.PredictedAt,
		UpdatedAt:   p.UpdatedAt,
		User: User{
			UserID:          p.UserID,
			UserDisplayName: p.UserDisplayName,
		},
		ResultType:           p.Result.Type,
		ResultPointsWon:      p.Result.PointsWon,
		ResultIsAcknowledged: p.Result.IsAcknowledged,
	}
}
