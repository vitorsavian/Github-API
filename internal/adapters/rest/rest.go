package rest

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	swaggerFiles "github.com/swaggo/files"
	ginSwag "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"sync"
	"time"
	"vitorsavian/github-api/docs"
	"vitorsavian/github-api/internal/adapters/controllers"
	"vitorsavian/github-api/internal/adapters/rest/health"

	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type Api struct {
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

func (h *Handler) NewApi() (*Api, error) {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(otelgin.Middleware("git-api"))

	base := router.Group("/git-api")

	docs.SwaggerInfo.BasePath = "/git-api"
	docs.SwaggerInfo.Title = "Git api"
	docs.SwaggerInfo.Description = "This is the git api documentation"
	docs.SwaggerInfo.Version = "1.0"
	fmt.Println(docs.SwaggerInfo.Host)

	fmt.Println(router.BasePath())
	base.GET("/documentation/*any", ginSwag.WrapHandler(swaggerFiles.Handler))
	base.GET("/health-check", health.HealthCheck)

	v1 := base.Group("/v1")

	gitRepositories := v1.Group("/repos")
	gitRepositories.GET("/:username", h.GitController.GetRepositoriesController)

	gitCommits := v1.Group("/commits")
	gitCommits.GET("/:username", h.GitController.GetCommitsController)

	server := &http.Server{Addr: fmt.Sprintf(":%d", h.Port), Handler: router}

	return &Api{
		server: server,
		port:   h.Port,
	}, nil
}

func (a *Api) Run() <-chan error {
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

func (a *Api) Shutdown() {
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
