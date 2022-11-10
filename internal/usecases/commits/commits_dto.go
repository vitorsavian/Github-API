package commits

type InputDto struct {
	UserName string
}

type OutputDto struct {
	Data RespData `json:"data"`
}

type RespData struct {
	Projects []ProjectInfo `json:"repos"`
}

type ProjectInfo struct {
	RepoName    string       `json:"name"`
	RepoOwner   string       `json:"owner"`
	UserCommits []UserCommit `json:"user-commits"`
}

type UserCommit struct {
	Url     string `json:"url"`
	Date    string `json:"date"`
	Message string `json:"message"`
}
