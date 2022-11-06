package github

import (
	"net/http"
	"time"
)

type GitService struct {
	httpClient http.Client
}

func NewGitService() *GitService {
	return &GitService{
		httpClient: http.Client{Timeout: time.Second * 30},
	}
}
