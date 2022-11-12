package repositories_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"vitorsavian/github-api/internal/adapters/services/git"
	gitMock "vitorsavian/github-api/internal/adapters/services/git/mock"
	"vitorsavian/github-api/internal/infrastructure/env"
	"vitorsavian/github-api/internal/usecases/repositories"
)

func TestRepos_GetRepos(t *testing.T) {
	input := repositories.InputDto{UserName: "vitorsavian"}
	mockGitService := gitMock.Service{}
	envs := env.Environment{GitServiceUrl: "git"}
	mockGitService.On("GetRepositoriesAscOrder", envs.GitServiceUrl, input.UserName).Return(
		[]git.GetRepositoriesResponse{
			{
				RepoId:      123123,
				RepoName:    "Admin-Catalog",
				RepoOwner:   "test",
				Description: "testing description",
				Forks:       1,
				OpenIssues:  1,
			},
			{
				RepoId:      233213,
				RepoName:    "cicd",
				RepoOwner:   "test",
				Description: "testing repos",
				Forks:       1,
				OpenIssues:  1,
			},
		}, nil,
	)

	useCase := repositories.New(&mockGitService, envs)

	expected := repositories.OutputDto{
		Data: []repositories.Repository{
			{
				RepoId:      123123,
				RepoName:    "Admin-Catalog",
				RepoOwner:   "test",
				Description: "testing description",
				Forks:       1,
				OpenIssues:  1,
			},
			{
				RepoId:      233213,
				RepoName:    "cicd",
				RepoOwner:   "test",
				Description: "testing repos",
				Forks:       1,
				OpenIssues:  1,
			},
		},
	}

	resp, err := useCase.GetRepositories(input)

	assert.Equal(t, nil, err, "error is not nil")
	assert.Equal(t, expected, resp, "response doenst match")

}
