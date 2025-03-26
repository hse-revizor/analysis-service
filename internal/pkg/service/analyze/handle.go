package analyze

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/4kayDev/logger/log"
	"github.com/google/uuid"
	"github.com/hse-revizor/analysis-service/internal/pkg/models"
	customjson "github.com/hse-revizor/analysis-service/internal/utils/json"
)

type RuleInput struct {
	RuleName    string `json:"rule_name"`
	FileContent string `json:"file_content"`
	Params      string `json:"params"`
}

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

	rule, err := s.rulesClient.GetRule(ctx, rulesetId.String())
	if err != nil {
		return nil, err
	}

	project, err := s.projectsClient.GetProject(ctx, projectId.String())
	if err != nil {
		return nil, err
	}

	parsedContent, err := s.parserClient.ParseProject(ctx, project.RepositoryURL)
	if err != nil {
		return nil, err
	}

	ruleInput := &RuleInput{
		RuleName:    rule.TypeId,
		FileContent: parsedContent.Content,
		Params:      rule.Params,
	}

	tmpFile, err := os.CreateTemp("", "rule-input-*.json")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file: %w", err)
	}
	defer os.Remove(tmpFile.Name())

	if err := customjson.WriteToFile(tmpFile, ruleInput); err != nil {
		return nil, fmt.Errorf("failed to write to temp file: %w", err)
	}
	tmpFile.Close()

	cmd := exec.Command("python", "strict_rules.py", tmpFile.Name())
	output, err := cmd.CombinedOutput()
	if err != nil {
		analyze.Status = "failed"
		analyze.RuleResults = []models.AnalyzeRuleResult{
			{
				Id:        uuid.New(),
				AnalyzeId: analyze.Id,
				RuleId:    rulesetId,
				Status:    "failed",
				Details:   fmt.Sprintf("Error running rule: %v", err),
			},
		}
		if updateErr := s.UpdateAnalyze(ctx, analyze); updateErr != nil {
			log.Error(updateErr, fmt.Sprintf("Failed to update failed analyze: %v", updateErr))
		}
		return nil, fmt.Errorf("failed to run python script: %w", err)
	}

	var scriptResult struct {
		Status  string `json:"status"`
		Details string `json:"details"`
	}
	if err := json.Unmarshal(output, &scriptResult); err != nil {
		return nil, fmt.Errorf("failed to parse script output: %w", err)
	}

	analyze.Status = "completed"
	analyze.RuleResults = []models.AnalyzeRuleResult{
		{
			Id:        uuid.New(),
			AnalyzeId: analyze.Id,
			RuleId:    rulesetId,
			Status:    scriptResult.Status,
			Details:   scriptResult.Details,
		},
	}
	analyze.FinishedAt = &now

	if err := s.UpdateAnalyze(ctx, analyze); err != nil {
		return nil, fmt.Errorf("failed to update analyze results: %w", err)
	}

	return analyze, nil
}

// @throws: ErrAnalyzeNotFound, ErrAnalyzeExists
func (s *Service) UpdateAnalyze(ctx context.Context, analyze *models.Analyze) error {
	err := s.storage.UpdateAnalyze(ctx, analyze)
	if err != nil {
		return err
	}
	log.Debug(fmt.Sprintf("Updated analyze: %s", customjson.ToColorJson(analyze)))
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
