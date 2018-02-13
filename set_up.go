package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"sync"
)

func getUserDir() string {
	user, err := user.Current()
	check(err)
	path := user.HomeDir
	return path
}

func setupDir() {
	homePath := getUserDir()
	path := homePath + FRONTEND_APPS_PATH
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("Creating dir")
		err := os.MkdirAll(path, 0755)
		check(err)		
	} else {
		fmt.Println("Dirs Created....already Exiting.")
		os.Exit(0)
	}
}

/*
Clones all the repos into  FRONTEND_APPS_PATH cds user into the dir of choice.
*/

func cloneRepos() {
	sliceLength := len(gitHubProjects)
	homePath := getUserDir()
	path := homePath + FRONTEND_APPS_PATH
	var wg sync.WaitGroup
	wg.Add(sliceLength)
	for i := 0; i < sliceLength; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println("Git Cloning ", gitHubProjects[i])
			cmd := exec.Command("git", "clone", gitHubProjects[i])
			cmd.Dir = path
			_, err := cmd.Output()
			check(err)
		}(i)
	}
	wg.Wait()
}

func setUp() {
	setupDir()	
	cloneRepos()
}
