package env

import (
	"os"
	"strconv"
	"strings"
)

//GetOrg - Gets Githhub Organization Namespace
func GetOrg() string {
	return os.Getenv("ORGANIZATION")
}

//GetAccessToken - Gets Access Token for Cloning Repos
func GetAccessToken() string {
	return os.Getenv("ACCESS_TOKEN")
}

//GetUserName - Gets UserName
func GetUserName() string {
	val := os.Getenv("GITHUB_USER_NAME")
	if val != "" {
		return val
	}
	return ""
}

//GetDestinationPath - Gets DestinationPath for dropping repos
func GetDestinationPath() string {
	return os.Getenv("DESTINATION_DIRECTORY")
}

//GetGitAPI - Get Organization API
func GetGitAPI() string {
	return os.Getenv("GITHUB_ORG_API")
}

//GetIgnoreList - Ignore Repos
func GetIgnoreList() []string {
	value := os.Getenv("IGNORE_REPOS")
	ignoreList := strings.Split(value, ",")
	return ignoreList
}

//CloneFork - Can clone forks
func CloneFork() bool {
	val := os.Getenv("IGNORE_FORK_REPOS")
	status, _ := strconv.ParseBool(val)
	return status
}
