package projects

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	requestbuidler "github.com/dr3dnought/request_builder"
	"github.com/hse-revizor/analysis-service/internal/pkg/clients/projects/models"
	"github.com/hse-revizor/analysis-service/internal/utils/config"
)

type Client struct {
	cfg     *config.Config
	builder *requestbuidler.RequestBuilder

	httpClient *http.Client
}

func New(cfg *config.Config) *Client {
	return &Client{
		cfg:     cfg,
		builder: requestbuidler.New(cfg.Client.ProjectsURL),

		httpClient: &http.Client{},
	}
}

func (c *Client) GetProject(ctx context.Context, projectID string) (*models.GetProjectDto, error) {
	response, err := c.builder.SetMethod("GET").SetPath("project/" + projectID).Build().Execute(c.httpClient)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	rawBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, parseApiError(response.StatusCode, rawBody)
	}

	project := new(models.GetProjectDto)
	err = json.Unmarshal(rawBody, project)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s can not be casted to ApiError type", string(rawBody)))
	}

	return project, nil
}
func parseApiError(statusCode int, data []byte) error {
	apiError := new(models.ApiError)
	err := json.Unmarshal(data, apiError)
	if err != nil {
		return errors.New(fmt.Sprintf("%s can not be casted to ApiError type", string(data)))
	}

	switch statusCode {
	case http.StatusBadRequest:
		return errors.New(apiError.Description)
	case http.StatusNotFound:
		return errors.New(apiError.Error)
	case http.StatusUnauthorized:
		return errors.New(apiError.Description)
	case http.StatusForbidden:
		return errors.New(apiError.Description)
	case http.StatusInternalServerError:
		return errors.New(apiError.Description)
	default:
		return errors.New(apiError.Description)
	}
}
