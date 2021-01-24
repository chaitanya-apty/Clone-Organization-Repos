package helpers

import (
	"io/ioutil"
	"log"
	"net/http"
)

//DownloadRepos - Todo... currently helper for downloading
func DownloadRepos(url string, token string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Cannot Download Repositories, Please try again with correct details")
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", "token "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode == 404 {
		log.Fatal("Cannot Download Repositories, Please try again with correct details")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}
