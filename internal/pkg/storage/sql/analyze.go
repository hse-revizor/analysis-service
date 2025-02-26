package sql

import (
	"context"

	"github.com/google/uuid"
	"github.com/hse-revizor/analysis-service/internal/pkg/models"
)

func (s *Storage) CreateAnalyze(ctx context.Context, analyze *models.Analyze) error {
	return s.db.WithContext(ctx).Create(analyze).Error
}

func (s *Storage) UpdateAnalyze(ctx context.Context, analyze *models.Analyze) error {
	return s.db.WithContext(ctx).Save(analyze).Error
}

func (s *Storage) GetAnalyzeById(ctx context.Context, id uuid.UUID) (*models.Analyze, error) {
	var analyze models.Analyze
	err := s.db.WithContext(ctx).Preload("RuleResults").First(&analyze, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &analyze, nil
}

func (s *Storage) GetAnalyzesByProjectId(ctx context.Context, projectId uuid.UUID) ([]models.Analyze, error) {
	var analyzes []models.Analyze
	err := s.db.WithContext(ctx).Where("project_id = ?", projectId).Find(&analyzes).Error
	return analyzes, err
}
