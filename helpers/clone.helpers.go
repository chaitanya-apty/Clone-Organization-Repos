package helpers

import "fmt"

//IgnoreRepo - Ignore Repos
func IgnoreRepo(currentRepo string, ignoreList []string) bool {
	if len(ignoreList) == 0 || ignoreList == nil {
		return false
	}
	for _, repo := range ignoreList {
		if repo == currentRepo {
			return true
		}
	}
	return false
}

//SkipForkedRepo - Ignore Repos
func SkipForkedRepo(isRepoFork bool, forkValue bool) bool {
	return isRepoFork && forkValue
}

//GetCloneCommand - Get Command String For Current Repo
func GetCloneCommand(repoName string, token string, userName string, orgName string) string {
	var baseCommand string
	if userName == "" {
		baseCommand = fmt.Sprintf("https://%s@github.com/%s/%s.git", token, orgName, repoName)
	}
	baseCommand = fmt.Sprintf("https://%s:%s@github.com/%s/%s.git", userName, token, orgName, repoName)
	return baseCommand
}
