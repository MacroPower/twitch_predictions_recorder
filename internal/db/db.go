package db

import (
	"fmt"

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
	GetSummary() ([]models.EventSummary, string, error)
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
	if err := gdb.setup(&event.PredictorState{}); err != nil {
		return fmt.Errorf("error migrating 'predictor_state': %w", err)
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
		return err
	}
	return nil
}

func (gdb *GormDB) GetSummary() ([]models.EventSummary, string, error) {
	var results []models.EventSummary
	var dbResults []struct {
		ID                      string
		ChannelName             string
		PredictionWindowSeconds int
		Title                   string
		Status                  string
		OutcomeColor            string
		OutcomeTitle            string
		OutcomeBadgeVersion     string
		OutcomeBadgeSetID       string
		OutcomeTotalPoints      int
		OutcomeTotalUsers       int
		OutcomeResultType       string
	}

	session := gdb.db //.Session(&gorm.Session{DryRun: true})

	tx := session.
		Select(
			"events.id AS id",
			"events.channel_name AS channel_name",
			"events.prediction_window_seconds AS prediction_window_seconds",
			"events.title AS title",
			"event_states.status AS status",
			"outcomes.color AS outcome_color",
			"outcomes.title AS outcome_title",
			"outcomes.badge_version AS outcome_badge_version",
			"outcomes.badge_set_id AS outcome_badge_set_id",
			"outcome_states.result_type AS outcome_result_type",
			"outcome_states.total_points AS outcome_total_points",
			"outcome_states.total_users AS outcome_total_users",
		).
		Table("events").
		Joins("JOIN (?) AS event_states ON events.id = event_states.event_id", session.
			Select("id", "event_id", "status", "max(timestamp)").
			Table("event_states").
			Group("event_id"),
		).
		Joins("JOIN outcomes ON event_states.id == outcomes.event_state_id").
		Joins("JOIN (?) AS outcome_states ON outcomes.id == outcome_states.outcome_id", session.
			Select("outcome_id", "total_points", "total_users", "result_type", "max(timestamp)").
			Table("outcome_states").
			Group("outcome_id"),
		).
		Find(&dbResults)
	if tx.Error != nil {
		return nil, "", tx.Error
	}

	rr := map[string]*models.EventSummary{}
	for _, r := range dbResults {
		if _, ok := rr[r.ID]; !ok {
			rr[r.ID] = &models.EventSummary{
				ID:                      r.ID,
				ChannelName:             r.ChannelName,
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
