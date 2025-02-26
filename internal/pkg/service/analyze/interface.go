package analyze

import (
	"context"

	"github.com/google/uuid"
	"github.com/hse-revizor/analysis-service/internal/pkg/models"
)

type storage interface {
	CreateAnalyze(context.Context, *models.Analyze) error
	GetAnalyzeById(ctx context.Context, id uuid.UUID) (*models.Analyze, error)
	UpdateAnalyze(context.Context, *models.Analyze) error
	GetAnalyzesByProjectId(ctx context.Context, projectId uuid.UUID) ([]models.Analyze, error)
}

type Service struct {
	storage storage
}

func New(storage storage) *Service {
	return &Service{storage: storage}
}
