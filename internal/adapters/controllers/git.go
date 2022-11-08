package controllers

import (
	"vitorsavian/github-api/internal/adapters/services/git"
	"vitorsavian/github-api/internal/infrastructure/env"

	"github.com/gin-gonic/gin"
)

type GitController struct {
	GitService git.Service
	Env        env.Environment
}

func NewGitController(gitService git.Service, env env.Environment) *GitController {
	return &GitController{
		GitService: gitService,
		Env:        env,
	}
}

func (g *GitController) GetRepositoriesController(c *gin.Context) {
	return
}

func (g *GitController) GetCommitsController(c *gin.Context) {
	return
}
