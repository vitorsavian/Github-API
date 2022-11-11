package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
	"vitorsavian/github-api/internal/adapters/controllers"
	"vitorsavian/github-api/internal/adapters/rest/health"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type api struct {
	server *http.Server
	port   int
}

type Handler struct {
	GitController controllers.GitController
	Port          int
}

func NewHandler(gitController controllers.GitController, port int) Handler {
	return Handler{
		GitController: gitController,
		Port:          port,
	}
}

func (h *Handler) NewApi() (*api, error) {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(otelgin.Middleware("git-api"))

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
		port:   h.Port,
	}, nil
}

func (a *api) Run() <-chan error {
	out := make(chan error)
	go func() {
		fmt.Printf("Server listening on port %d\n", a.port)
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println(err)
			out <- errors.Wrap(err, "failed to listen and serve api")
		}
	}()
	return out
}

func (a *api) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := a.server.Shutdown(ctx)
		if err != nil {
			log.Println("Server forced to shutdown")
		}
	}()
}
