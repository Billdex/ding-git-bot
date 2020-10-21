package webhook

// 简易Github消息结构
type GithubMsg struct {
	Repository githubRepository `json:"repository"`
	Sender     githubSender     `json:"sender"`
}

// 仓库信息
type githubRepository struct {
	Id               int64                 `json:"id"`
	NodeId           string                `json:"node_id"`
	Name             string                `json:"name"`
	FullName         string                `json:"full_name"`
	Private          bool                  `json:"private"`
	Owner            githubRepositoryOwner `json:"owner"`
	HtmlUrl          string                `json:"html_url"`
	Description      string                `json:"description"`
	Fork             bool                  `json:"fork"`
	Url              string                `json:"url"`
	ForksUrl         string                `json:"forks_url"`
	KeysUrl          string                `json:"keys_url"`
	CollaboratorsUrl string                `json:"collaborators_url"`
	TeamsUrl         string                `json:"teams_url"`
	HooksUrl         string                `json:"hooks_url"`
	IssueEventsUrl   string                `json:"issue_events_url"`
	EventsUrl        string                `json:"events_url"`
	AssigneesUrl     string                `json:"assignees_url"`
	BranchesUrl      string                `json:"branches_url"`
	TagsUrl          string                `json:"tags_url"`
	BlobsUrl         string                `json:"blobs_url"`
	GitTagsUrl       string                `json:"git_tags_url"`
	GitRefsUrl       string                `json:"git_refs_url"`
	TreesUrl         string                `json:"trees_url"`
	StatusesUrl      string                `json:"statuses_url"`
	LanguagesUrl     string                `json:"languages_url"`
	StargazersUrl    string                `json:"stargazers_url"`
	ContributorsUrl  string                `json:"contributors_url"`
	SubscribersUrl   string                `json:"subscribers_url"`
	SubscriptionUrl  string                `json:"subscription_url"`
	CommitsUrl       string                `json:"commits_url"`
	GitCommitsUrl    string                `json:"git_commits_url"`
	CommentsUrl      string                `json:"comments_url"`
	IssueCommentUrl  string                `json:"issue_comment_url"`
	ContentsUrl      string                `json:"contents_url"`
	CompareUrl       string                `json:"compare_url"`
	MergesUrl        string                `json:"merges_url"`
	ArchiveUrl       string                `json:"archive_url"`
	DownloadsUrl     string                `json:"downloads_url"`
	IssuesUrl        string                `json:"issues_url"`
	PullsUrl         string                `json:"pulls_url"`
	MilestonesUrl    string                `json:"milestones_url"`
	NotificationsUrl string                `json:"notifications_url"`
	LabelsUrl        string                `json:"labels_url"`
	ReleasesUrl      string                `json:"releases_url"`
	DeploymentsUrl   string                `json:"deployments_url"`
	CreatedAt        interface{}           `json:"created_at"`
	UpdatedAt        interface{}           `json:"updated_at"`
	PushedAt         interface{}           `json:"pushed_at"`
	GitUrl           string                `json:"git_url"`
	SSHUrl           string                `json:"ssh_url"`
	CloneUrl         string                `json:"clone_url"`
	SVNUrl           string                `json:"svn_url"`
	Homepage         string                `json:"homepage"`
	Size             int64                 `json:"size"`
	StargazersCount  int64                 `json:"stargazers_count"`
	WatchersCount    int64                 `json:"watchers_count"`
	Language         string                `json:"language"`
	HasIssues        bool                  `json:"has_issues"`
	HasProjects      bool                  `json:"has_projects"`
	HasDownloads     bool                  `json:"has_downloads"`
	HasWiki          bool                  `json:"has_wiki"`
	HasPages         bool                  `json:"has_pages"`
	ForksCount       int64                 `json:"forks_count"`
	MirrorUrl        string                `json:"mirror_url"`
	Archived         bool                  `json:"archived"`
	Disabled         bool                  `json:"disabled"`
	OpenIssuesCount  int64                 `json:"open_issues_count"`
	License          string                `json:"license"`
	Forks            int64                 `json:"forks"`
	OpenIssues       int64                 `json:"open_issues"`
	Watchers         int64                 `json:"watchers"`
	DefaultBranch    string                `json:"default_branch"`
	Stargazers       int64                 `json:"stargazers"`
	MasterBranch     string                `json:"master_branch"`
}

// 仓库所有者信息
type githubRepositoryOwner struct {
	Login             string `json:"login"`
	Id                int64  `json:"id"`
	NodeId            string `json:"node_id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"following_url"`
	GistsUrl          string `json:"gists_url"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	ReposUrl          string `json:"repos_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

// Webhooks触发者
type githubSender struct {
	Login             string `json:"login"`
	Id                int64  `json:"id"`
	NodeId            string `json:"node_id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"Gravatar_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"following_url"`
	GistsUrl          string `json:"gists_url"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	ReposUrl          string `json:"repos_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

// Ping消息结构
type GithubPingMsg struct {
}
