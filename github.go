package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"net/http"
	"os"
)

var GH_BASE_URL string = "https://api.github.com/"

func spinUpPr(opts RepoOpts) {
	fmt.Println("Spinning up PR for " + opts.path)
	access_token := os.Getenv("GITHUB_TOKEN")
	path := "repos"
	owner := "hyfn"
	repo := opts.path
	resource := "pulls"
	url := GH_BASE_URL + path + "/" + owner + "/" + repo + "/" + resource + "?access_token=" + access_token
	payload := map[string]string{
		"title": opts.branch,
		"body":  "Please pull this in for " + opts.branch,
		"head":  opts.branch,
		"base":  "master"}

	json, err := json.Marshal(payload)
	check(err)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	req.SetBasicAuth(access_token, "x-oauth-basic")
	client := http.Client{}
	check(err)

	res, err := client.Do(req)
	check(err)

	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body)) 
}
