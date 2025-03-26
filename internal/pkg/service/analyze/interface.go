package analyze

import (
	"context"

	"github.com/google/uuid"
	"github.com/hse-revizor/analysis-service/internal/pkg/clients/parser"
	"github.com/hse-revizor/analysis-service/internal/pkg/clients/projects"
	"github.com/hse-revizor/analysis-service/internal/pkg/clients/rules"
	"github.com/hse-revizor/analysis-service/internal/pkg/models"
)

type storage interface {
	CreateAnalyze(context.Context, *models.Analyze) error
	GetAnalyzeById(ctx context.Context, id uuid.UUID) (*models.Analyze, error)
	UpdateAnalyze(context.Context, *models.Analyze) error
	GetAnalyzesByProjectId(ctx context.Context, projectId uuid.UUID) ([]models.Analyze, error)
}

type Service struct {
	rulesClient    *rules.Client
	projectsClient *projects.Client
	parserClient   *parser.Client
	storage        storage
}

func New(storage storage, rulesClient *rules.Client, projectsClient *projects.Client, parserClient *parser.Client) *Service {
	return &Service{storage: storage, rulesClient: rulesClient, projectsClient: projectsClient, parserClient: parserClient}
}
