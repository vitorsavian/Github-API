package commits_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"vitorsavian/github-api/internal/adapters/services/git"
	gitMock "vitorsavian/github-api/internal/adapters/services/git/mock"
	"vitorsavian/github-api/internal/infrastructure/env"
	"vitorsavian/github-api/internal/usecases/commits"
)

func TestCommits_GetCommitsAndReposWithOneRepoReceived(t *testing.T) {
	input := commits.InputDto{UserName: "vitorsavian"}
	mockGitService := gitMock.Service{}
	envs := env.Environment{GitServiceUrl: "git"}
	mockGitService.On("GetRecentRepositoriesPushed", envs.GitServiceUrl, input.UserName).Return(
		[]git.GetRepositoriesResponse{
			{
				RepoId:      123123,
				RepoName:    "test",
				RepoOwner:   "testing",
				Description: "testing description",
				Forks:       1,
				OpenIssues:  1,
			},
		}, nil,
	)

	mockGitService.On("GetCommitsPerRepository", envs.GitServiceUrl, "testing", input.UserName, "test").Return(
		git.GetCommitsRepositoryResponse{
			User: git.CommitsUser{
				Name:  "testing",
				Email: "testing@test.com",
			},
			RepoOwner: "testing",
			RepoName:  "test",
			LastCommits: []git.Commit{
				{
					Date:      "2022-01-01",
					Message:   "testing",
					CommitUrl: "dale.com.br",
				},
				{
					Date:      "2022-01-02",
					Message:   "testing something",
					CommitUrl: "dale.com.br",
				},
			},
		}, nil,
	)

	useCase := commits.New(&mockGitService, envs)

	expected := commits.OutputDto{
		Data: commits.RespData{
			Projects: []commits.ProjectInfo{
				{
					RepoName:  "test",
					RepoOwner: "testing",
					UserCommits: []commits.UserCommit{
						{
							Url:     "dale.com.br",
							Date:    "2022-01-01",
							Message: "testing",
						},
						{
							Date:    "2022-01-02",
							Message: "testing something",
							Url:     "dale.com.br",
						},
					},
				},
			},
		}}

	resp, err := useCase.GetCommits(input)

	assert.Equal(t, nil, err, "error is not nil")
	assert.Equal(t, expected, resp, "response doenst match")

}

func TestCommits_GetCommitsAndReposWithoutRepoReceived(t *testing.T) {
	input := commits.InputDto{UserName: "vitorsavian"}
	mockGitService := gitMock.Service{}
	envs := env.Environment{GitServiceUrl: "git"}
	mockGitService.On("GetRecentRepositoriesPushed", envs.GitServiceUrl, input.UserName).Return(
		[]git.GetRepositoriesResponse{}, nil,
	)

	useCase := commits.New(&mockGitService, envs)

	expected := commits.OutputDto{
		Data: commits.RespData{
			Projects: []commits.ProjectInfo{},
		}}

	resp, err := useCase.GetCommits(input)

	assert.Equal(t, nil, err, "error is not nil")
	assert.Equal(t, expected, resp, "response doenst match")

}

func TestCommits_GetCommitsAndReposWithMoreThanOneRepoReceived(t *testing.T) {
	input := commits.InputDto{UserName: "vitorsavian"}
	mockGitService := gitMock.Service{}
	envs := env.Environment{GitServiceUrl: "git"}
	mockGitService.On("GetRecentRepositoriesPushed", envs.GitServiceUrl, input.UserName).Return(
		[]git.GetRepositoriesResponse{
			{
				RepoId:      123123,
				RepoName:    "test",
				RepoOwner:   "testing",
				Description: "testing description",
				Forks:       1,
				OpenIssues:  1,
			},
			{
				RepoId:      123321,
				RepoName:    "test2",
				RepoOwner:   "testing2",
				Description: "testing description2",
				Forks:       2,
				OpenIssues:  2,
			},
		}, nil,
	)

	mockGitService.On("GetCommitsPerRepository", envs.GitServiceUrl, "testing", input.UserName, "test").Return(
		git.GetCommitsRepositoryResponse{
			User: git.CommitsUser{
				Name:  "testing",
				Email: "testing@test.com",
			},
			RepoOwner: "testing",
			RepoName:  "test",
			LastCommits: []git.Commit{
				{
					Date:      "2022-01-01",
					Message:   "testing",
					CommitUrl: "dale.com.br",
				},
				{
					Date:      "2022-01-02",
					Message:   "testing something",
					CommitUrl: "dale.com.br",
				},
			},
		}, nil,
	)

	mockGitService.On("GetCommitsPerRepository", envs.GitServiceUrl, "testing2", input.UserName, "test2").Return(
		git.GetCommitsRepositoryResponse{
			User: git.CommitsUser{
				Name:  "testing2",
				Email: "testing2@test.com",
			},
			RepoOwner: "testing2",
			RepoName:  "test2",
			LastCommits: []git.Commit{
				{
					Date:      "2022-01-01",
					Message:   "testing2",
					CommitUrl: "dale.com.br",
				},
				{
					Date:      "2022-01-02",
					Message:   "testing something2",
					CommitUrl: "dale.com.br",
				},
			},
		}, nil,
	)

	useCase := commits.New(&mockGitService, envs)

	expected := commits.OutputDto{
		Data: commits.RespData{
			Projects: []commits.ProjectInfo{
				{
					RepoName:  "test",
					RepoOwner: "testing",
					UserCommits: []commits.UserCommit{
						{
							Url:     "dale.com.br",
							Date:    "2022-01-01",
							Message: "testing",
						},
						{
							Date:    "2022-01-02",
							Message: "testing something",
							Url:     "dale.com.br",
						},
					},
				},
				{
					RepoName:  "test2",
					RepoOwner: "testing2",
					UserCommits: []commits.UserCommit{
						{
							Url:     "dale.com.br",
							Date:    "2022-01-01",
							Message: "testing2",
						},
						{
							Date:    "2022-01-02",
							Message: "testing something2",
							Url:     "dale.com.br",
						},
					},
				},
			},
		}}

	resp, err := useCase.GetCommits(input)

	assert.Equal(t, nil, err, "error is not nil")
	assert.Equal(t, expected, resp, "response doenst match")
}
