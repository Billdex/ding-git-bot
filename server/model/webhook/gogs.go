package webhook

import "time"

//gogs消息结构与github很相似

type GogsMsg struct {
	Repository gogsRepository `json:"repository"`
	Sender     gogsSender     `json:"sender"`
}

type gogsRepository struct {
	Id              int64               `json:"id"`
	Owner           gogsRepositoryOwner `json:"owner"`
	Name            string              `json:"name"`
	FullName        string              `json:"full_name"`
	Description     string              `json:"description"`
	Private         bool                `json:"private"`
	Fork            bool                `json:"fork"`
	HtmlUrl         string              `json:"html_url"`
	SSHUrl          string              `json:"ssh_url"`
	CloneUrl        string              `json:"clone_url"`
	Website         string              `json:"website"`
	StarsCount      int64               `json:"stars_count"`
	ForksCount      int64               `json:"forks_count"`
	WatchersCount   int64               `json:"watchers_count"`
	OpenIssuesCount int64               `json:"open_issues_count"`
	DefaultBranch   string              `json:"default_branch"`
	CreatedAt       time.Time           `json:"created_at"`
	UpdatedAt       time.Time           `json:"updated_at"`
}

type gogsRepositoryOwner struct {
	Id        int64  `json:"id"`
	Login     string `json:"login"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	AvatarUrl string `json:"avatar_url"`
	Username  string `json:"username"`
}

type gogsPusher struct {
	Id        int64  `json:"id"`
	Login     string `json:"login"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	AvatarUrl string `json:"avatar_url"`
	Username  string `json:"username"`
}

type gogsSender struct {
	Id        int64  `json:"id"`
	Login     string `json:"login"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	AvatarUrl string `json:"avatar_url"`
	Username  string `json:"username"`
}

type gogsCommit struct {
	Id      string `json:"id"`
	Message string `json:"message"`
	Url     string `json:"url"`
	Author  struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Username string `json:"username"`
	} `json:"author"`
	Committer struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Username string `json:"username"`
	} `json:"committer"`
	Timestamp time.Time `json:"timestamp"`
}

type GogsPushMsg struct {
	Ref        string         `json:"ref"`
	Before     string         `json:"before"`
	After      string         `json:"after"`
	CompareUrl string         `json:"compare_url"`
	Commits    []gogsCommit   `json:"commits"`
	Repository gogsRepository `json:"repository"`
	Pusher     gogsPusher     `json:"pusher"`
	Sender     gogsSender     `json:"sender"`
}
