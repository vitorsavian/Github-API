package git

type Service interface {
	GetRepositories(user string) (GetRepositoriesResponse, error)
}

type GetCommitsRepositoryResponse struct {
	Commits CommitsResponse
}

type CommitsResponse struct {
	User CommitsUser
}

type CommitsUser struct {
	Name  string
	Email string
}

type GetRepositoriesResponse struct {
	RepoName    string
	RepoId      string
	Description string
	Forks       int
	OpenIssues  int
}
