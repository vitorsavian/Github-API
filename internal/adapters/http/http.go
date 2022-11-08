package http

import (
	"fmt"
	"net/http"
	"vitorsavian/github-api/internal/adapters/controllers"
	"vitorsavian/github-api/internal/adapters/http/health"

	"github.com/gin-gonic/gin"
)

type api struct {
	server *http.Server
}

type Handler struct {
	GitController controllers.GitController
	Port          int
}

func (h *Handler) NewApi() (*api, error) {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Recovery())

	base := router.Group("/git-api")
	base.GET("/health-check", health.HealthCheck)

	v1 := base.Group("/v1")

	gitRepositories := v1.Group("/repos")
	gitRepositories.GET("/:username", h.GitController.GetRepositoriesController)

	gitCommits := v1.Group("/commits")
	gitCommits.GET("/:username", h.GitController.GetCommitsController)

	server := &http.Server{Addr: fmt.Sprintf(":%d", h.Port), Handler: router}

	return &api{
		server: server,
	}, nil
}
