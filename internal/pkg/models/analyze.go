package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Analyze struct {
	Id          uuid.UUID           `gorm:"primaryKey;column:id"`
	ProjectId   uuid.UUID           `gorm:"column:project_id"`
	Status      string              `gorm:"column:status"`
	RulesetId   uuid.UUID           `gorm:"column:ruleset_id"`
	RuleResults []AnalyzeRuleResult `gorm:"foreignKey:AnalyzeId"`

	StartedAt  *time.Time `gorm:"column:started_at"`
	FinishedAt *time.Time `gorm:"column:finished_at"`
	CreatedAt  *time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  *time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

type AnalyzeRuleResult struct {
	Id        uuid.UUID `gorm:"primaryKey;column:id"`
	AnalyzeId uuid.UUID `gorm:"column:analyze_id"`
	RuleId    uuid.UUID `gorm:"column:rule_id"`
	Status    string    `gorm:"column:status"` // passed, failed, skipped
	Details   string    `gorm:"column:details;type:text"`

	CreatedAt *time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (g *Analyze) BeforeCreate(tx *gorm.DB) error {
	if g.Id == uuid.Nil {
		g.Id = uuid.New()
	}
	return nil
}
