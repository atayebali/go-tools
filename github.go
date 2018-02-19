package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var GH_BASE_URL string = "https://api.github.com/"

func spinUpPr(opts RepoOpts) {
	if len(os.Getenv("GITHUB_TOKEN")) == 0 {
		fmt.Println("Git hub token not found as ENV var.")
		fmt.Println("A Pr will not be generated")
		os.Exit(0)
	}

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

	code := res.StatusCode
	fmt.Printf("API Returned HTTP STATUS CODE: %d\n ", code)
}
