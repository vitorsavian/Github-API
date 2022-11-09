package repositories

type InputDto struct {
	UserName string
}

type OutputDto struct {
	Data []Repository `json:"data"`
}

type Repository struct {
	RepoId      int
	RepoName    string
	RepoOwner   string
	Description string
	Forks       int
	OpenIssues  int
}
