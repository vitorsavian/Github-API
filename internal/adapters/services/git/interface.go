package git

type Service interface {
	GetRepositoriesAscOrder(urlService, user string) ([]GetRepositoriesResponse, error)
	GetRecentRepositoriesPushed(urlService, user string) ([]GetRepositoriesResponse, error)
	GetCommitsPerRepository(urlService, owner, user, project string) (GetCommitsRepositoryResponse, error)
}

type GetCommitsRepositoryResponse struct {
	User        CommitsUser
	RepoOwner   string
	RepoName    string
	LastCommits []Commit
}

type CommitsUser struct {
	Name  string
	Email string
}

type Commit struct {
	Date      string
	Message   string
	CommitUrl string
}

type GetRepositoriesResponse struct {
	RepoId      int
	RepoName    string
	RepoOwner   string
	Description string
	Forks       int
	OpenIssues  int
}
