package commits

// import (
// 	"vitorsavian/github-api/internal/adapters/services/git"
// 	"vitorsavian/github-api/internal/domain/value"
// 	"vitorsavian/github-api/internal/infrastructure/env"
// )

// type GetCommits struct {
// 	GitService git.Service
// 	Env        env.Environment
// }

// var (
// 	ErrEmptyUserName = value.NewCustomError("err.get.commits.username_empty", "The username is empty", "The informed username is empty")
// )

// func New(gitService git.Service, envs env.Environment) *GetCommits {
// 	return &GetCommits{
// 		GitService: gitService,
// 		Env:        envs,
// 	}
// }

// func (g *GetCommits) GetCommits(input InputDto) (OutputDto, error) {

// }
