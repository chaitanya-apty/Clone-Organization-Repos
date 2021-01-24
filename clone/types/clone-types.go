package types

//DownloadRepos - Interface for VC
type DownloadRepos interface {
	ValidateOrg() bool
	ValidateDestination() bool
	CloneReposToDestination()
}

//GithubReposInfo - Struct
type GithubReposInfo struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Fork          bool   `json:"fork"`
	CloneURL      string `json:"clone_url"`
	DefaultBranch string `json:"default_branch"`
	Owner         struct {
		Name string `json:"login"`
	} `json:"owner"`
}
