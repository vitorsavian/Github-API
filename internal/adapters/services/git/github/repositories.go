package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"vitorsavian/github-api/internal/adapters/services/git"
)

var (
	accept = "application/vnd.github+json"
)

func (g *GitService) GetRepositoriesAscOrder(urlService, user string) ([]git.GetRepositoriesResponse, error) {
	url := fmt.Sprintf("%s/users/%s/repos?type=owner&per_page=100", urlService, user)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []git.GetRepositoriesResponse{}, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")

	res, err := g.httpClient.Do(req)
	if err != nil {
		return []git.GetRepositoriesResponse{}, err
	}

	defer res.Body.Close()

	var responseContract []struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Owner struct {
			Login string `json:"login"`
		} `json:"owner"`
		Description string `json:"description"`
		Forks       int    `json:"forks_count"`
		OpenIssues  int    `json:"open_issues_count"`
	}

	var errorResponseContract struct {
		Message string `json:"message"`
	}

	if res.StatusCode != http.StatusOK {
		if err = json.NewDecoder(res.Body).Decode(&errorResponseContract); err != nil {
			return []git.GetRepositoriesResponse{}, err
		}
		return []git.GetRepositoriesResponse{}, errors.New(errorResponseContract.Message)
	}

	if err = json.NewDecoder(res.Body).Decode(&responseContract); err != nil {
		return []git.GetRepositoriesResponse{}, err
	}

	if len(responseContract) == 0 {
		return []git.GetRepositoriesResponse{}, err
	}

	var repoResponse []git.GetRepositoriesResponse

	for _, value := range responseContract {
		repoResponse = append(repoResponse, git.GetRepositoriesResponse{
			RepoId:      value.Id,
			RepoName:    value.Name,
			RepoOwner:   value.Owner.Login,
			Description: value.Description,
			Forks:       value.Forks,
			OpenIssues:  value.OpenIssues,
		})
	}

	return repoResponse, nil
}
