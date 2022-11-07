package github

import "vitorsavian/github-api/internal/adapters/services/git"

func (g *GitService) GetRepositories(user string) (git.GetRepositoriesResponse, error) {
	return git.GetRepositoriesResponse{}, nil
}
