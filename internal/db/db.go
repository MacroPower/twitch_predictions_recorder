package db

import (
	"fmt"
	"strings"
	"time"

	"github.com/MacroPower/twitch_predictions_recorder/internal/api/models"
	"github.com/MacroPower/twitch_predictions_recorder/internal/event"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GormDB struct {
	db *gorm.DB
}

type DB interface {
	SetupDefaults() error
	AddEvents(...event.Event) error
	GetSummary(string) ([]models.EventSummary, string, error)
	GetDetails(string) ([]models.EventDetails, string, error)
}

func (gdb *GormDB) SetupDefaults() error {
	if err := gdb.setup(&event.Event{}); err != nil {
		return fmt.Errorf("error migrating 'event': %w", err)
	}
	if err := gdb.setup(&event.EventState{}); err != nil {
		return fmt.Errorf("error migrating 'event_state': %w", err)
	}
	if err := gdb.setup(&event.Outcome{}); err != nil {
		return fmt.Errorf("error migrating 'outcome': %w", err)
	}
	if err := gdb.setup(&event.OutcomeState{}); err != nil {
		return fmt.Errorf("error migrating 'outcome_state': %w", err)
	}
	if err := gdb.setup(&event.Predictor{}); err != nil {
		return fmt.Errorf("error migrating 'predictor': %w", err)
	}
	if err := gdb.setup(&event.User{}); err != nil {
		return fmt.Errorf("error migrating 'user': %w", err)
	}

	return nil
}

func (gdb *GormDB) AddEvents(events ...event.Event) error {
	tx := gdb.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(events)

	return tx.Error
}

func (gdb *GormDB) setup(obj interface{}) error {
	if err := gdb.db.AutoMigrate(obj); err != nil {
		return fmt.Errorf("auto migration error: %w", err)
	}

	return nil
}

func (gdb *GormDB) GetSummary(eventID string) ([]models.EventSummary, string, error) {
	var results []models.EventSummary
	var dbResults []struct {
		ID                      string
		ChannelName             string
		CreatedAt               time.Time
		PredictionWindowSeconds int
		Title                   string
		Timestamp               string
		Status                  string
		OutcomeColor            string
		OutcomeTitle            string
		OutcomeBadgeVersion     string
		OutcomeBadgeSetID       string
		OutcomeTotalPoints      int
		OutcomeTotalUsers       int
		OutcomeResultType       string
	}

	session := gdb.db // .Session(&gorm.Session{DryRun: true})

	query := session.
		Select(
			"events.id AS id",
			"events.channel_name AS channel_name",
			"events.created_at AS created_at",
			"events.prediction_window_seconds AS prediction_window_seconds",
			"events.title AS title",
			"event_states.timestamp AS timestamp",
			"event_states.status AS status",
			"outcomes.color AS outcome_color",
			"outcomes.title AS outcome_title",
			"outcomes.badge_version AS outcome_badge_version",
			"outcomes.badge_set_id AS outcome_badge_set_id",
			"outcomes.result_type AS outcome_result_type",
			"outcomes.total_points AS outcome_total_points",
			"outcomes.total_users AS outcome_total_users",
		).
		Table("events").
		Joins("JOIN (?) AS event_states ON events.id = event_states.event_id", session.
			Select("id", "event_id", "status", "max(timestamp) AS timestamp").
			Table("event_states").
			Group("event_id"),
		).
		Joins("JOIN (?) AS outcomes ON events.id = outcomes.event_id", session.
			Select(
				"outcome_id",
				"event_id",
				"color",
				"title",
				"badge_version",
				"badge_set_id",
				"total_points",
				"total_users",
				"result_type",
				"max(timestamp)",
			).
			Table("outcomes").
			Joins("JOIN outcome_states ON outcomes.id = outcome_states.outcome_id").
			Group("outcome_id"),
		)

	if eventID != "" {
		query = query.Where("events.id = ?", eventID)
	}

	tx := query.Find(&dbResults)
	if tx.Error != nil {
		return nil, "", tx.Error
	}

	rr := map[string]*models.EventSummary{}
	for _, r := range dbResults {
		if _, ok := rr[r.ID]; !ok {
			ts, err := parseTimestamp(r.Timestamp)
			if err != nil {
				return nil, "", err
			}
			rr[r.ID] = &models.EventSummary{
				ID:                      r.ID,
				Timestamp:               ts,
				ChannelName:             r.ChannelName,
				CreatedAt:               r.CreatedAt,
				PredictionWindowSeconds: r.PredictionWindowSeconds,
				Title:                   r.Title,
				Status:                  r.Status,
				Outcomes:                []models.OutcomeSummary{},
			}
		}
		rr[r.ID].Outcomes = append(rr[r.ID].Outcomes, models.OutcomeSummary{
			Color:        r.OutcomeColor,
			Title:        r.OutcomeTitle,
			BadgeVersion: r.OutcomeBadgeVersion,
			BadgeSetID:   r.OutcomeBadgeSetID,
			TotalPoints:  r.OutcomeTotalPoints,
			TotalUsers:   r.OutcomeTotalUsers,
			ResultType:   r.OutcomeResultType,
		})
	}
	for _, v := range rr {
		results = append(results, *v)
	}

	return results, tx.Statement.SQL.String(), nil
}

type getDetailsDBRow struct {
	ID                      string
	ChannelName             string
	CreatedAt               time.Time
	PredictionWindowSeconds int
	Title                   string
	EventStateID            string
	Timestamp               time.Time
	Status                  string
	OutcomeID               string
	OutcomeColor            string
	OutcomeTitle            string
	OutcomeBadgeVersion     string
	OutcomeBadgeSetID       string
	OutcomeStateID          string
	OutcomeTotalPoints      int
	OutcomeTotalUsers       int
	OutcomeResultType       string
	PredictorID             string
	PredictorName           string
	PredictorPoints         int
}

func (gdb *GormDB) GetDetails(eventID string) ([]models.EventDetails, string, error) {
	var dbResults []getDetailsDBRow

	session := gdb.db // .Session(&gorm.Session{DryRun: true})

	query := session.
		Select(
			"events.id AS id",
			"events.channel_name AS channel_name",
			"events.created_at AS created_at",
			"events.prediction_window_seconds AS prediction_window_seconds",
			"events.title AS title",
			"event_states.id AS event_state_id",
			"event_states.timestamp AS timestamp",
			"event_states.status AS status",
			"outcomes.id AS outcome_id",
			"outcomes.color AS outcome_color",
			"outcomes.title AS outcome_title",
			"outcomes.badge_version AS outcome_badge_version",
			"outcomes.badge_set_id AS outcome_badge_set_id",
			"outcome_states.id AS outcome_state_id",
			"outcome_states.result_type AS outcome_result_type",
			"outcome_states.total_points AS outcome_total_points",
			"outcome_states.total_users AS outcome_total_users",
			"predictors.predicted_by_id AS predictor_id",
			"users.user_display_name AS predictor_name",
			"predictors.points AS predictor_points",
		).
		Table("events").
		Joins("JOIN event_states ON events.id = event_states.event_id").
		Joins("JOIN outcome_states ON outcome_states.event_state_id = event_states.id").
		Joins("JOIN outcomes ON outcomes.id = outcome_states.outcome_id").
		Joins("JOIN predictors ON outcome_states.id = predictors.outcome_state_id").
		Joins("JOIN users ON predictors.predicted_by_id = users.user_id").
		Where("events.id = ?", eventID)

	tx := query.Find(&dbResults)
	sql := tx.Statement.SQL.String()
	if tx.Error != nil {
		return nil, sql, tx.Error
	}

	results := transformGetDetailsDBResults(dbResults)

	return results, sql, nil
}

func parseTimestamp(ts string) (time.Time, error) {
	t, err := time.Parse("2006-01-02 15:04:05.999999999Z07:00", ts)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse timestamp: %w", err)
	}

	return t, nil
}

func transformGetDetailsDBResults(dbResults []getDetailsDBRow) []models.EventDetails {
	eventDetailsMap := map[string]*models.EventDetails{}
	eventSeriesMap := map[string]*models.EventSeries{}
	outcomeDetailsMap := map[string]*models.OutcomeDetails{}

	for _, result := range dbResults {
		// Create or get the event detail
		if _, ok := eventDetailsMap[result.ID]; !ok {
			eventDetailsMap[result.ID] = &models.EventDetails{
				ID:                      result.ID,
				ChannelName:             result.ChannelName,
				CreatedAt:               result.CreatedAt,
				PredictionWindowSeconds: result.PredictionWindowSeconds,
				Title:                   result.Title,
				EventSeries:             []models.EventSeries{},
			}
		}

		// Create or get the event series
		seriesKey := result.ID + result.EventStateID
		if _, ok := eventSeriesMap[seriesKey]; !ok {
			eventSeriesMap[seriesKey] = &models.EventSeries{
				ID:        result.EventStateID,
				Timestamp: result.Timestamp,
				Status:    result.Status,
				Outcomes:  []models.OutcomeDetails{},
			}
		}

		// Create or get the outcome detail
		outcomeKey := seriesKey + result.OutcomeStateID
		if _, ok := outcomeDetailsMap[outcomeKey]; !ok {
			outcomeDetailsMap[outcomeKey] = &models.OutcomeDetails{
				ID:            result.OutcomeStateID,
				Color:         result.OutcomeColor,
				Title:         result.OutcomeTitle,
				BadgeVersion:  result.OutcomeBadgeVersion,
				BadgeSetID:    result.OutcomeBadgeSetID,
				Timestamp:     result.Timestamp,
				TotalPoints:   result.OutcomeTotalPoints,
				TotalUsers:    result.OutcomeTotalUsers,
				ResultType:    result.OutcomeResultType,
				TopPredictors: []models.Predictor{},
			}
		}

		// Add the predictor to the outcome detail
		if result.PredictorID != "" {
			predictor := models.Predictor{
				User: models.User{
					UserID:          result.PredictorID,
					UserDisplayName: result.PredictorName,
				},
				Points: result.PredictorPoints,
			}
			outcomeDetailsMap[outcomeKey].TopPredictors = append(outcomeDetailsMap[outcomeKey].TopPredictors, predictor)
		}
	}

	// Add event series to event details and outcomes to event series
	for _, eventDetail := range eventDetailsMap {
		for seriesKey, eventSeries := range eventSeriesMap {
			if strings.HasPrefix(seriesKey, eventDetail.ID) {
				for outcomeKey, outcomeDetail := range outcomeDetailsMap {
					if strings.HasPrefix(outcomeKey, seriesKey) {
						eventSeries.Outcomes = append(eventSeries.Outcomes, *outcomeDetail)
					}
				}
				eventDetail.EventSeries = append(eventDetail.EventSeries, *eventSeries)
			}
		}
	}

	eventDetailsSlice := make([]models.EventDetails, 0, len(eventDetailsMap))
	for _, v := range eventDetailsMap {
		eventDetailsSlice = append(eventDetailsSlice, *v)
	}

	return eventDetailsSlice
}
