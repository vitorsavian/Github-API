package repositories

import (
	"vitorsavian/github-api/internal/adapters/services/git"
	"vitorsavian/github-api/internal/domain/value"
	"vitorsavian/github-api/internal/infrastructure/env"
)

type Repositories struct {
	GitService git.Service
	Env        env.Environment
}

var (
	ErrEmptyUserName = value.NewCustomError("err.get.repositories.username_empty", "The username is empty", "The informed username is empty")
)

func New(gitService git.Service, envs env.Environment) *Repositories {
	return &Repositories{
		GitService: gitService,
		Env:        envs,
	}
}

func (r *Repositories) GetRepositories(input InputDto) (OutputDto, error) {
	if input.UserName == "" {
		return OutputDto{}, ErrEmptyUserName
	}

	repoResponse, err := r.GitService.GetRepositoriesAscOrder(r.Env.GitServiceUrl, input.UserName)
	if err != nil {
		return OutputDto{}, err
	}

	if len(repoResponse) == 0 {
		return OutputDto{
			Data: []Repository{},
		}, nil
	}

	output := OutputDto{
		Data: make([]Repository, len(repoResponse)),
	}

	for index, value := range repoResponse {
		outputItem := Repository{
			RepoId:      value.RepoId,
			RepoName:    value.RepoName,
			RepoOwner:   value.RepoOwner,
			Description: value.Description,
			Forks:       value.Forks,
			OpenIssues:  value.OpenIssues,
		}
		output.Data[index] = outputItem
	}

	return output, nil
}
