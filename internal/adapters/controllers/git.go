package controllers

import (
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

func (g *GitController) GetRepositoriesController(c *gin.Context) {
	log.Println("GetRepositoriesController Called")

	inputDto := repositories.InputDto{
		UserName: c.Param("username"),
	}

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
