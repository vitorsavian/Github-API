package git

type Service interface {
	GetRepositoriesAscOrder(urlService, user string) ([]GetRepositoriesResponse, error)
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
	RepoId      int
	RepoName    string
	RepoOwner   string
	Description string
	Forks       int
	OpenIssues  int
}
