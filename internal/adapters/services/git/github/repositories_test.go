package github_test

import (
	"testing"
	"vitorsavian/github-api/internal/infrastructure/env"
)

func configTest() *env.Environment {
	return env.GetEnvironment()
}

func TestGitService_GetRepositories(t *testing.T) {
	config := configTest()

}
