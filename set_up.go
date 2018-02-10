package main

import (
	"fmt"
	"os"
	"sync"
	"os/user"
)

const FRONTEND_APPS_PATH = "/FRONT_END_APPS"

var gitHubProjects = []string{
	"git@github.com:rtyley/small-test-repo.git",
	"git@github.com:kelseyhightower/nocode.git"}

func getUserDir() string {
	user, err := user.Current()
	check(err)
	path := user.HomeDir
	return path
}

func setupDir(){
	path := getUserDir()
	fmt.Println(path + FRONTEND_APPS_PATH)
	if _, err := os.Stat(path + FRONTEND_APPS_PATH); os.IsNotExist(err) {
		fmt.Println("Creating dir")
		err := os.MkdirAll(path + FRONTEND_APPS_PATH, 0755)
		check(err)

	}
	fmt.Println("Setup is complete....")
}

func cloneRepos() {
	sliceLength := len(gitHubProjects)
	var wg sync.WaitGroup
	wg.Add(sliceLength)
	for i := 0; i < sliceLength; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(gitHubProjects[i])
		}(i)
	}
	wg.Wait()
}



func setUp() {
	setupDir()
	cloneRepos()
}
