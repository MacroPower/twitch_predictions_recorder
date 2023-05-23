package db

import (
	"testing"
	"time"

	"github.com/MacroPower/twitch_predictions_recorder/internal/api/models"
	"github.com/stretchr/testify/require"
)

func TestTransform(t *testing.T) {
	now := time.Now()

	dbResults := []getDetailsDBRow{
		{
			ID:                      "1",
			ChannelName:             "Channel 1",
			CreatedAt:               now,
			PredictionWindowSeconds: 10,
			Title:                   "Title 1",
			EventStateID:            "EventState 1",
			Timestamp:               now,
			Status:                  "Status 1",
			OutcomeID:               "Outcome 1",
			OutcomeColor:            "Red",
			OutcomeTitle:            "Outcome Title 1",
			OutcomeBadgeVersion:     "1",
			OutcomeBadgeSetID:       "BadgeSet 1",
			OutcomeStateID:          "State 1",
			OutcomeTotalPoints:      100,
			OutcomeTotalUsers:       10,
			OutcomeResultType:       "ResultType 1",
			PredictorID:             "Predictor 1",
			PredictorName:           "Predictor Name 1",
			PredictorPoints:         50,
		},
		{
			ID:                      "1",
			ChannelName:             "Channel 1",
			CreatedAt:               now,
			PredictionWindowSeconds: 10,
			Title:                   "Title 1",
			EventStateID:            "EventState 1",
			Timestamp:               now,
			Status:                  "Status 1",
			OutcomeID:               "Outcome 1",
			OutcomeColor:            "Red",
			OutcomeTitle:            "Outcome Title 1",
			OutcomeBadgeVersion:     "1",
			OutcomeBadgeSetID:       "BadgeSet 1",
			OutcomeStateID:          "State 1",
			OutcomeTotalPoints:      100,
			OutcomeTotalUsers:       10,
			OutcomeResultType:       "ResultType 1",
			PredictorID:             "Predictor 2",
			PredictorName:           "Predictor Name 2",
			PredictorPoints:         500,
		},
		{
			ID:                      "1",
			ChannelName:             "Channel 1",
			CreatedAt:               now,
			PredictionWindowSeconds: 10,
			Title:                   "Title 1",
			EventStateID:            "EventState 1",
			Timestamp:               now,
			Status:                  "Status 1",
			OutcomeID:               "Outcome 2",
			OutcomeColor:            "Red",
			OutcomeTitle:            "Outcome Title 2",
			OutcomeBadgeVersion:     "2",
			OutcomeBadgeSetID:       "BadgeSet 2",
			OutcomeStateID:          "State 2",
			OutcomeTotalPoints:      100,
			OutcomeTotalUsers:       10,
			OutcomeResultType:       "ResultType 2",
			PredictorID:             "Predictor 3",
			PredictorName:           "Predictor Name 3",
			PredictorPoints:         500,
		},
	}

	expected := []models.EventDetails{
		{
			ID:                      "1",
			ChannelName:             "Channel 1",
			CreatedAt:               now,
			PredictionWindowSeconds: 10,
			Title:                   "Title 1",
			EventSeries: []models.EventSeries{
				{
					ID:        "EventState 1",
					Timestamp: now,
					Status:    "Status 1",
					Outcomes: []models.OutcomeDetails{
						{
							ID:           "State 1",
							Color:        "Red",
							Title:        "Outcome Title 1",
							BadgeVersion: "1",
							BadgeSetID:   "BadgeSet 1",
							Timestamp:    now,
							TotalPoints:  100,
							TotalUsers:   10,
							ResultType:   "ResultType 1",
							TopPredictors: []models.Predictor{
								{
									User: models.User{
										UserID:          "Predictor 1",
										UserDisplayName: "Predictor Name 1",
									},
									Points: 50,
								},
								{
									User: models.User{
										UserID:          "Predictor 2",
										UserDisplayName: "Predictor Name 2",
									},
									Points: 500,
								},
							},
						},
						{
							ID:           "State 2",
							Color:        "Red",
							Title:        "Outcome Title 2",
							BadgeVersion: "2",
							BadgeSetID:   "BadgeSet 2",
							Timestamp:    now,
							TotalPoints:  100,
							TotalUsers:   10,
							ResultType:   "ResultType 2",
							TopPredictors: []models.Predictor{
								{
									User: models.User{
										UserID:          "Predictor 3",
										UserDisplayName: "Predictor Name 3",
									},
									Points: 500,
								},
							},
						},
					},
				},
			},
		},
	}

	actual := transformGetDetailsDbResults(dbResults)

	require.Equal(t, expected, actual)
}
