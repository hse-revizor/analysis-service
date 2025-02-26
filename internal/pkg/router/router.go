package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	docs "github.com/hse-revizor/analysis-service/docs"
	"github.com/hse-revizor/analysis-service/internal/pkg/service/analyze"
	"github.com/hse-revizor/analysis-service/internal/utils/config"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	cfg     *config.Config
	service *analyze.Service
}

func NewRouter(cfg *config.Config, service *analyze.Service) *Handler {
	return &Handler{
		cfg:     cfg,
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	api := gin.New()

	api.Use(gin.Recovery())
	api.Use(gin.Logger())
	api.Use(cors.Default())

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	docs.SwaggerInfo.Title = "Analysis Service API"
	docs.SwaggerInfo.Description = "API Documentation"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8787"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http"}

	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router := api.Group("/api")
	{
		analyzes := router.Group("/analyze")
		{
			analyzes.POST("", h.CreateAnalyze)
			analyzes.GET("/:id", h.GetAnalyze)
		}

		projects := router.Group("/projects")
		{
			projects.GET("/:project_id/analyzes", h.GetProjectAnalyzes)
		}
	}

	return api
}
