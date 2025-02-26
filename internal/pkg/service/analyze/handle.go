package analyze

import (
	"context"
	"fmt"
	"time"

	"github.com/4kayDev/logger/log"
	"github.com/google/uuid"
	"github.com/hse-revizor/analysis-service/internal/pkg/models"
	"github.com/hse-revizor/analysis-service/internal/utils/json"
)

// @throws: ErrAnalyzeNotFound, ErrAnalyzeExists
func (s *Service) CreateAnalyze(ctx context.Context, projectId, rulesetId uuid.UUID) (*models.Analyze, error) {
	now := time.Now()
	analyze := &models.Analyze{
		ProjectId: projectId,
		RulesetId: rulesetId,
		Status:    "pending",
		StartedAt: &now,
		CreatedAt: &now,
	}

	if err := s.storage.CreateAnalyze(ctx, analyze); err != nil {
		return nil, err
	}

	return analyze, nil
}

// @throws: ErrAnalyzeNotFound, ErrAnalyzeExists
func (s *Service) UpdateAnalyze(ctx context.Context, analyze *models.Analyze) error {
	err := s.storage.UpdateAnalyze(ctx, analyze)
	if err != nil {
		return err
	}
	log.Debug(fmt.Sprintf("Updated analyze: %s", json.ToColorJson(analyze)))
	return nil
}

func (s *Service) GetAnalyze(ctx context.Context, id uuid.UUID) (*models.Analyze, error) {
	return s.storage.GetAnalyzeById(ctx, id)
}

func (s *Service) UpdateAnalyzeStatus(ctx context.Context, id uuid.UUID, status string, results []models.AnalyzeRuleResult) error {
	analyze, err := s.storage.GetAnalyzeById(ctx, id)
	if err != nil {
		return err
	}

	now := time.Now()
	analyze.Status = status
	analyze.RuleResults = results
	analyze.FinishedAt = &now

	return s.storage.UpdateAnalyze(ctx, analyze)
}

func (s *Service) GetAnalyzesByProjectId(ctx context.Context, projectId uuid.UUID) ([]models.Analyze, error) {
	return s.storage.GetAnalyzesByProjectId(ctx, projectId)
}
