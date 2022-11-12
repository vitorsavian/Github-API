package commits

import (
	"vitorsavian/github-api/internal/adapters/services/git"
	"vitorsavian/github-api/internal/domain/value"
	"vitorsavian/github-api/internal/infrastructure/env"
)

type Commits struct {
	GitService git.Service
	Env        env.Environment
}

var (
	ErrEmptyUserName = value.NewCustomError("err.get.commits.username_empty", "The username is empty", "The informed username is empty")
)

func New(gitService git.Service, envs env.Environment) *Commits {
	return &Commits{
		GitService: gitService,
		Env:        envs,
	}
}

func (g *Commits) GetCommits(input InputDto) (OutputDto, error) {
	if input.UserName == "" {
		return OutputDto{}, ErrEmptyUserName
	}

	repoResponse, err := g.GitService.GetRecentRepositoriesPushed(g.Env.GitServiceUrl, input.UserName)
	if err != nil {
		return OutputDto{}, err
	}

	if len(repoResponse) == 0 {
		return OutputDto{
			Data: RespData{
				Projects: []ProjectInfo{},
			},
		}, nil
	}

	var commitResponses []git.GetCommitsRepositoryResponse

	for _, rep := range repoResponse {
		commitResponse, _ := g.GitService.GetCommitsPerRepository(g.Env.GitServiceUrl, rep.RepoOwner, input.UserName, rep.RepoName)
		commitResponses = append(commitResponses, commitResponse)
	}

	var outputProjectInfo []ProjectInfo

	for _, project := range commitResponses {
		var commitsUser []UserCommit
		for _, commit := range project.LastCommits {
			commitsUser = append(commitsUser, UserCommit{
				Url:     commit.CommitUrl,
				Date:    commit.Date,
				Message: commit.Message,
			})
		}

		projectInfo := ProjectInfo{
			RepoName:    project.RepoName,
			RepoOwner:   project.RepoOwner,
			UserCommits: commitsUser,
		}

		outputProjectInfo = append(outputProjectInfo, projectInfo)
	}

	output := OutputDto{
		Data: RespData{
			Projects: outputProjectInfo,
		},
	}

	return output, nil
}
