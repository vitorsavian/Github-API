package controllers

import (
	"fmt"
	"log"
	"net/http"
	"vitorsavian/github-api/internal/adapters/services/git"
	"vitorsavian/github-api/internal/infrastructure/env"
	"vitorsavian/github-api/internal/usecases/commits"
	"vitorsavian/github-api/internal/usecases/repositories"

	"github.com/gin-gonic/gin"
)

type GitController struct {
	GitService git.Service
	Env        *env.Environment
}

func NewGitController(gitService git.Service, envs *env.Environment) *GitController {
	return &GitController{
		GitService: gitService,
		Env:        envs,
	}
}

// @BasePath /v1

// Get Repos godoc
// @Summary Get Repos
// @Description Route to get all repos in alphabetical order
// @Tags Repo
// @Accept json
// @Produce json
// @Param  username path string true "username from the user"
// @Success 200
// @Router /v1/repos/{username} [get]
func (g *GitController) GetRepositoriesController(c *gin.Context) {
	log.Println("GetRepositoriesController Called")

	inputDto := repositories.InputDto{
		UserName: c.Param("username"),
	}
	fmt.Println(inputDto.UserName)
	useCase := repositories.New(g.GitService, *g.Env)

	resp, err := useCase.GetRepositories(inputDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @title Swagger Example API
// @BasePath /v1/
func (g *GitController) GetCommitsController(c *gin.Context) {
	log.Println("GetCommitsController Called")

	inputDto := commits.InputDto{
		UserName: c.Param("username"),
	}

	useCase := commits.New(g.GitService, *g.Env)

	resp, err := useCase.GetCommits(inputDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
