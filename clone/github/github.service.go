package github

import (
	ty "clone-repos/clone/types"
	env "clone-repos/env"
	hp "clone-repos/helpers"
	"fmt"
	"os/exec"

	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//Service - GithubService
type Service struct {
	ReposURL    string `json:"repos_url"`
	Login       string `json:"login"`
	Destination string
}

//ValidateOrg - Validate Organization,token
func (svc *Service) ValidateOrg() bool {
	api := env.GetGitAPI()
	org := env.GetOrg()
	url := api + org
	res, err := http.Get(url)
	if err != nil || res.StatusCode == 404 {
		log.Fatal("Organization Not Found, Please update .env with valid values")
		return false
	}
	if res.StatusCode == 200 {
		log.Printf("Organization Validated")
	}
	data, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(data, &svc)
	if err != nil {
		log.Fatal("Cannot Retrieve Details From Your ORG")
	}
	res.Body.Close()
	return true
}

//ValidateDestination - Validate DestinationPath
func (svc *Service) ValidateDestination() bool {
	destDir := env.GetDestinationPath()
	_, err := os.Stat(destDir)
	if os.IsNotExist(err) {
		log.Fatal("Destination Directory Not Found")
		return false
	}
	log.Println("Destination Directory Validated")
	svc.Destination = destDir
	return true
}

func logOutput(name string) {
	fmt.Println("________________________________________________________")
	fmt.Println("\nCloning in Progress <===>", name, ",please wait..")
}

//CloneReposToDestination - Download Repos to Destination
func (svc *Service) CloneReposToDestination() {
	token := env.GetAccessToken()
	ignoreList := env.GetIgnoreList()
	cloneFork := env.CloneFork()
	userName := env.GetUserName()

	skipped := 0
	repos := []ty.GithubReposInfo{}
	reposInfo := hp.DownloadRepos(svc.ReposURL, token)
	err := json.Unmarshal(reposInfo, &repos)
	if err != nil {
		log.Fatal("Repositories Information Error")
	}
	fmt.Println("\n\n<<<< Process Might Take Time, Grab a coffee! :P >>>>")
	for _, r := range repos {
		if hp.IgnoreRepo(r.Name, ignoreList) || hp.SkipForkedRepo(r.Fork, cloneFork) {
			skipped++
			fmt.Println("\n*****Ignored Repo:", r.Name, "*******")
			continue
		}
		prepCloneCommand := hp.GetCloneCommand(r.Name, token, userName, r.Owner.Name)
		cmd := exec.Command("git", "clone", prepCloneCommand)
		cmd.Dir = svc.Destination
		if err := cmd.Start(); err != nil {
			log.Printf("Failed to execute Git Command: %v", err)
			return
		}

		logOutput(r.Name)

		// Wait Before Moving to Next clone:
		if err := cmd.Wait(); err != nil {
			log.Printf("Could not clone: %s, %v", r.Name, err)
		}

		log.Println("Completed For: [", r.ID, "|", r.Name, "|", r.DefaultBranch, "]")
	}
	log.Println("<< Cloning Completed Successfully, Check your Destination Directory >>")
}
