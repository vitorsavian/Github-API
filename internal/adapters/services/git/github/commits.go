package github

import "vitorsavian/github-api/internal/adapters/services/git"

func (g *GitService) GetCommitsPerRepository(user, project string) (git.GetRepositoriesResponse, error) {
	return git.GetRepositoriesResponse{}, nil
}
