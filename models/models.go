package models

import (
	"time"
)

// PRResponse represents the top-level response from GitHub PR API
type PRResponse struct {
	TotalCount        int      `json:"total_count"`
	IncompleteResults bool     `json:"incomplete_results"`
	Items             []PRItem `json:"items"`
}

// PRItem represents a single pull request item
type PRItem struct {
	URL                   string      `json:"url"`
	RepositoryURL         string      `json:"repository_url"`
	LabelsURL             string      `json:"labels_url"`
	CommentsURL           string      `json:"comments_url"`
	EventsURL             string      `json:"events_url"`
	HTMLURL               string      `json:"html_url"`
	ID                    int64       `json:"id"`
	NodeID                string      `json:"node_id"`
	Number                int         `json:"number"`
	Title                 string      `json:"title"`
	User                  User        `json:"user"`
	Labels                []Label     `json:"labels"`
	State                 string      `json:"state"`
	Locked                bool        `json:"locked"`
	Assignee              *User       `json:"assignee"`
	Assignees             []User      `json:"assignees"`
	Milestone             *Milestone  `json:"milestone"`
	Comments              int         `json:"comments"`
	CreatedAt             time.Time   `json:"created_at"`
	UpdatedAt             time.Time   `json:"updated_at"`
	ClosedAt              *time.Time  `json:"closed_at"`
	AuthorAssociation     string      `json:"author_association"`
	ActiveLockReason      *string     `json:"active_lock_reason"`
	Draft                 bool        `json:"draft"`
	PullRequest           PullRequest `json:"pull_request"`
	Body                  string      `json:"body"`
	Reactions             Reactions   `json:"reactions"`
	TimelineURL           string      `json:"timeline_url"`
	PerformedViaGithubApp *string     `json:"performed_via_github_app"`
	StateReason           *string     `json:"state_reason"`
	Score                 float64     `json:"score"`
}

// User represents a GitHub user
type User struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	UserViewType      string `json:"user_view_type"`
	SiteAdmin         bool   `json:"site_admin"`
}

// Label represents a GitHub issue/PR label
type Label struct {
	ID          int64  `json:"id"`
	NodeID      string `json:"node_id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Default     bool   `json:"default"`
}

// Milestone represents a GitHub milestone
type Milestone struct {
	URL          string     `json:"url"`
	HTMLURL      string     `json:"html_url"`
	LabelsURL    string     `json:"labels_url"`
	ID           int64      `json:"id"`
	NodeID       string     `json:"node_id"`
	Number       int        `json:"number"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Creator      User       `json:"creator"`
	OpenIssues   int        `json:"open_issues"`
	ClosedIssues int        `json:"closed_issues"`
	State        string     `json:"state"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DueOn        *time.Time `json:"due_on"`
	ClosedAt     *time.Time `json:"closed_at"`
}

// PullRequest represents the pull request specific fields
type PullRequest struct {
	URL      string     `json:"url"`
	HTMLURL  string     `json:"html_url"`
	DiffURL  string     `json:"diff_url"`
	PatchURL string     `json:"patch_url"`
	MergedAt *time.Time `json:"merged_at"`
}

// Reactions represents reaction counts on a PR
type Reactions struct {
	URL        string `json:"url"`
	TotalCount int    `json:"total_count"`
	PlusOne    int    `json:"+1"`
	MinusOne   int    `json:"-1"`
	Laugh      int    `json:"laugh"`
	Hooray     int    `json:"hooray"`
	Confused   int    `json:"confused"`
	Heart      int    `json:"heart"`
	Rocket     int    `json:"rocket"`
	Eyes       int    `json:"eyes"`
}

// Links represents the _links field in GitHub PR API
type Links struct {
	Self           Link `json:"self"`
	HTML           Link `json:"html"`
	Issue          Link `json:"issue"`
	Comments       Link `json:"comments"`
	ReviewComments Link `json:"review_comments"`
	ReviewComment  Link `json:"review_comment"`
	Commits        Link `json:"commits"`
	Statuses       Link `json:"statuses"`
}

// Link represents a single link object in GitHub API
type Link struct {
	Href string `json:"href"`
}

// Repository represents a GitHub repository
type Repository struct {
	ID               int64     `json:"id"`
	NodeID           string    `json:"node_id"`
	Name             string    `json:"name"`
	FullName         string    `json:"full_name"`
	Private          bool      `json:"private"`
	Owner            User      `json:"owner"`
	HTMLURL          string    `json:"html_url"`
	Description      *string   `json:"description"`
	Fork             bool      `json:"fork"`
	URL              string    `json:"url"`
	ForksURL         string    `json:"forks_url"`
	KeysURL          string    `json:"keys_url"`
	CollaboratorsURL string    `json:"collaborators_url"`
	TeamsURL         string    `json:"teams_url"`
	HooksURL         string    `json:"hooks_url"`
	IssueEventsURL   string    `json:"issue_events_url"`
	EventsURL        string    `json:"events_url"`
	AssigneesURL     string    `json:"assignees_url"`
	BranchesURL      string    `json:"branches_url"`
	TagsURL          string    `json:"tags_url"`
	BlobsURL         string    `json:"blobs_url"`
	GitTagsURL       string    `json:"git_tags_url"`
	GitRefsURL       string    `json:"git_refs_url"`
	TreesURL         string    `json:"trees_url"`
	StatusesURL      string    `json:"statuses_url"`
	LanguagesURL     string    `json:"languages_url"`
	StargazersURL    string    `json:"stargazers_url"`
	ContributorsURL  string    `json:"contributors_url"`
	SubscribersURL   string    `json:"subscribers_url"`
	SubscriptionURL  string    `json:"subscription_url"`
	CommitsURL       string    `json:"commits_url"`
	GitCommitsURL    string    `json:"git_commits_url"`
	CommentsURL      string    `json:"comments_url"`
	IssueCommentURL  string    `json:"issue_comment_url"`
	ContentsURL      string    `json:"contents_url"`
	CompareURL       string    `json:"compare_url"`
	MergesURL        string    `json:"merges_url"`
	ArchiveURL       string    `json:"archive_url"`
	DownloadsURL     string    `json:"downloads_url"`
	IssuesURL        string    `json:"issues_url"`
	PullsURL         string    `json:"pulls_url"`
	MilestonesURL    string    `json:"milestones_url"`
	NotificationsURL string    `json:"notifications_url"`
	LabelsURL        string    `json:"labels_url"`
	ReleasesURL      string    `json:"releases_url"`
	DeploymentsURL   string    `json:"deployments_url"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	PushedAt         time.Time `json:"pushed_at"`
	GitURL           string    `json:"git_url"`
	SSHURL           string    `json:"ssh_url"`
	CloneURL         string    `json:"clone_url"`
	SvnURL           string    `json:"svn_url"`
	Homepage         *string   `json:"homepage"`
	Size             int       `json:"size"`
	StargazersCount  int       `json:"stargazers_count"`
	WatchersCount    int       `json:"watchers_count"`
	Language         string    `json:"language"`
	HasIssues        bool      `json:"has_issues"`
	HasProjects      bool      `json:"has_projects"`
	HasDownloads     bool      `json:"has_downloads"`
	HasWiki          bool      `json:"has_wiki"`
	HasPages         bool      `json:"has_pages"`
	HasDiscussions   bool      `json:"has_discussions"`
	ForksCount       int       `json:"forks_count"`
	MirrorURL        *string   `json:"mirror_url"`
	Archived         bool      `json:"archived"`
	Disabled         bool      `json:"disabled"`
	OpenIssuesCount  int       `json:"open_issues_count"`
	License          *string   `json:"license"`
	AllowForking     bool      `json:"allow_forking"`
	IsTemplate       bool      `json:"is_template"`
	WebCommitSignoff bool      `json:"web_commit_signoff_required"`
	Topics           []string  `json:"topics"`
	Visibility       string    `json:"visibility"`
	Forks            int       `json:"forks"`
	OpenIssues       int       `json:"open_issues"`
	Watchers         int       `json:"watchers"`
	DefaultBranch    string    `json:"default_branch"`
}

// CommitRef represents a Git reference in a PR
type CommitRef struct {
	Label string     `json:"label"`
	Ref   string     `json:"ref"`
	SHA   string     `json:"sha"`
	User  User       `json:"user"`
	Repo  Repository `json:"repo"`
}

// PullRequestDetail represents detailed information about a GitHub pull request
type PullRequestDetail struct {
	URL                 string     `json:"url"`
	ID                  int64      `json:"id"`
	NodeID              string     `json:"node_id"`
	HTMLURL             string     `json:"html_url"`
	DiffURL             string     `json:"diff_url"`
	PatchURL            string     `json:"patch_url"`
	IssueURL            string     `json:"issue_url"`
	Number              int        `json:"number"`
	State               string     `json:"state"`
	Locked              bool       `json:"locked"`
	Title               string     `json:"title"`
	User                User       `json:"user"`
	Body                *string    `json:"body"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
	ClosedAt            *time.Time `json:"closed_at"`
	MergedAt            *time.Time `json:"merged_at"`
	MergeCommitSHA      string     `json:"merge_commit_sha"`
	Assignee            *User      `json:"assignee"`
	Assignees           []User     `json:"assignees"`
	RequestedReviewers  []User     `json:"requested_reviewers"`
	RequestedTeams      []Team     `json:"requested_teams"`
	Labels              []Label    `json:"labels"`
	Milestone           *Milestone `json:"milestone"`
	Draft               bool       `json:"draft"`
	CommitsURL          string     `json:"commits_url"`
	ReviewCommentsURL   string     `json:"review_comments_url"`
	ReviewCommentURL    string     `json:"review_comment_url"`
	CommentsURL         string     `json:"comments_url"`
	StatusesURL         string     `json:"statuses_url"`
	Head                CommitRef  `json:"head"`
	Base                CommitRef  `json:"base"`
	Links               Links      `json:"_links"`
	AuthorAssociation   string     `json:"author_association"`
	AutoMerge           *string    `json:"auto_merge"`
	ActiveLockReason    *string    `json:"active_lock_reason"`
	Merged              bool       `json:"merged"`
	Mergeable           *bool      `json:"mergeable"`
	Rebaseable          *bool      `json:"rebaseable"`
	MergeableState      string     `json:"mergeable_state"`
	MergedBy            *User      `json:"merged_by"`
	Comments            int        `json:"comments"`
	ReviewComments      int        `json:"review_comments"`
	MaintainerCanModify bool       `json:"maintainer_can_modify"`
	Commits             int        `json:"commits"`
	Additions           int        `json:"additions"`
	Deletions           int        `json:"deletions"`
	ChangedFiles        int        `json:"changed_files"`
}

// Team represents a GitHub team
type Team struct {
	ID              int64  `json:"id"`
	NodeID          string `json:"node_id"`
	URL             string `json:"url"`
	HTMLURL         string `json:"html_url"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	Description     string `json:"description"`
	Privacy         string `json:"privacy"`
	Permission      string `json:"permission"`
	MembersURL      string `json:"members_url"`
	RepositoriesURL string `json:"repositories_url"`
}

type PRReview struct {
	ID                int64     `json:"id"`
	NodeID            string    `json:"node_id"`
	User              User      `json:"user"`
	Body              string    `json:"body"`
	State             string    `json:"state"`
	HTMLURL           string    `json:"html_url"`
	PullRequestURL    string    `json:"pull_request_url"`
	AuthorAssociation string    `json:"author_association"`
	Links             Links     `json:"_links"`
	SubmittedAt       time.Time `json:"submitted_at"`
	CommitID          string    `json:"commit_id"`
}

type ReviewLinks struct {
	HTML        Link `json:"html"`
	PullRequest Link `json:"pull_request"`
}
