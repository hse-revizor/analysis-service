package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	analyze "github.com/hse-revizor/analysis-service/internal/pkg/router/dto"
)

// @Summary Create new analyze
// @Tags Analyze
// @Accept json
// @Produce json
// @Param request body analyze.CreateAnalyzeRequest true "Analyze creation request"
// @Success 200 {object} analyze.AnalyzeResponse
// @Router /api/analyze [post]
func (h *Handler) CreateAnalyze(c *gin.Context) {
	var req analyze.CreateAnalyzeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.service.CreateAnalyze(c.Request.Context(), req.ProjectId, req.RulesetId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, analyze.ToAnalyzeResponse(result))
}

// @Summary Get analyze by ID
// @Tags Analyze
// @Produce json
// @Param id path string true "Analyze ID"
// @Success 200 {object} analyze.AnalyzeResponse
// @Router /api/analyze/{id} [get]
func (h *Handler) GetAnalyze(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid analyze id"})
		return
	}

	result, err := h.service.GetAnalyze(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, analyze.ToAnalyzeResponse(result))
}

// @Summary Get analyzes by project ID
// @Tags Analyze
// @Produce json
// @Param project_id path string true "Project ID"
// @Success 200 {array} analyze.AnalyzeResponse
// @Router /api/projects/{project_id}/analyzes [get]
func (h *Handler) GetProjectAnalyzes(c *gin.Context) {
	projectId, err := uuid.Parse(c.Param("project_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project id"})
		return
	}

	results, err := h.service.GetAnalyzesByProjectId(c.Request.Context(), projectId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := make([]*analyze.AnalyzeResponse, len(results))
	for i, result := range results {
		response[i] = analyze.ToAnalyzeResponse(&result)
	}

	c.JSON(http.StatusOK, response)
}
