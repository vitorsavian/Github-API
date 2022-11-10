package github

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"vitorsavian/github-api/internal/adapters/services/git"
)

func (g *GitService) GetCommitsPerRepository(urlService, owner, user, project string) (git.GetCommitsRepositoryResponse, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/commits?author=%s&per_page=5", urlService, owner, project, user)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return git.GetCommitsRepositoryResponse{}, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")

	res, err := g.httpClient.Do(req)
	if err != nil {
		return git.GetCommitsRepositoryResponse{}, err
	}

	defer res.Body.Close()

	var responseContract []struct {
		Commit struct {
			Url    string `json:"url"`
			Author struct {
				Name  string
				Email string
				Date  string
			} `json:"author"`
			Message string `json:"message"`
		} `json:"commit"`
	}

	var errorResponseContract struct {
		Message string `json:"message"`
	}

	if res.StatusCode != http.StatusOK {
		if err = json.NewDecoder(res.Body).Decode(&errorResponseContract); err != nil {
			return git.GetCommitsRepositoryResponse{}, err
		}
		return git.GetCommitsRepositoryResponse{}, errors.New(errorResponseContract.Message)
	}

	if err = json.NewDecoder(res.Body).Decode(&responseContract); err != nil {
		return git.GetCommitsRepositoryResponse{}, err
	}

	if len(responseContract) == 0 {
		return git.GetCommitsRepositoryResponse{}, nil
	}

	var commits []git.Commit

	for _, value := range responseContract {
		commits = append(commits, git.Commit{
			Date:      value.Commit.Author.Date,
			Message:   value.Commit.Message,
			CommitUrl: value.Commit.Url,
		})
	}

	commitResponse := git.GetCommitsRepositoryResponse{
		User: git.CommitsUser{
			Name:  responseContract[0].Commit.Author.Name,
			Email: responseContract[0].Commit.Author.Email,
		},
		RepoOwner:   owner,
		RepoName:    project,
		LastCommits: commits,
	}

	return commitResponse, nil
}
