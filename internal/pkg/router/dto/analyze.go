package analyze

import (
	"time"

	"github.com/google/uuid"
	"github.com/hse-revizor/analysis-service/internal/pkg/models"
)

type CreateAnalyzeRequest struct {
	ProjectId uuid.UUID `json:"project_id" binding:"required"`
	RulesetId uuid.UUID `json:"ruleset_id" binding:"required"`
}

type AnalyzeResponse struct {
	Id          uuid.UUID               `json:"id"`
	ProjectId   uuid.UUID               `json:"project_id"`
	Status      string                  `json:"status"`
	RulesetId   uuid.UUID               `json:"ruleset_id"`
	RuleResults []AnalyzeResultResponse `json:"rule_results,omitempty"`
	StartedAt   *time.Time              `json:"started_at,omitempty"`
	FinishedAt  *time.Time              `json:"finished_at,omitempty"`
	CreatedAt   *time.Time              `json:"created_at"`
	UpdatedAt   *time.Time              `json:"updated_at"`
}

type AnalyzeResultResponse struct {
	Id        uuid.UUID  `json:"id"`
	RuleId    uuid.UUID  `json:"rule_id"`
	Status    string     `json:"status"`
	Details   string     `json:"details"`
	CreatedAt *time.Time `json:"created_at"`
}

func ToAnalyzeResponse(analyze *models.Analyze) *AnalyzeResponse {
	response := &AnalyzeResponse{
		Id:         analyze.Id,
		ProjectId:  analyze.ProjectId,
		Status:     analyze.Status,
		RulesetId:  analyze.RulesetId,
		StartedAt:  analyze.StartedAt,
		FinishedAt: analyze.FinishedAt,
		CreatedAt:  analyze.CreatedAt,
		UpdatedAt:  analyze.UpdatedAt,
	}

	if len(analyze.RuleResults) > 0 {
		response.RuleResults = make([]AnalyzeResultResponse, len(analyze.RuleResults))
		for i, result := range analyze.RuleResults {
			response.RuleResults[i] = AnalyzeResultResponse{
				Id:        result.Id,
				RuleId:    result.RuleId,
				Status:    result.Status,
				Details:   result.Details,
				CreatedAt: result.CreatedAt,
			}
		}
	}

	return response
}
