package main

import (
	"os/exec"
	"fmt"
)

func gitCommit(message string, path string) {
	gitAdd := exec.Command("git", "add", ".")
	gitAdd.Dir = path

	_, err := gitAdd.Output()
	check(err)

	gitCommit := exec.Command("git", "commit", "-am", message)
	gitCommit.Dir = path


	fmt.Println("Commiting inside" + path)
	_, err1 := gitCommit.Output()
	check(err1)
}

func gitPush(opts RepoOpts) {
	cmd := exec.Command("git", "push", "origin", opts.branch)
	cmd.Dir = opts.path
	fmt.Println("Pushing to github: " + opts.branch)
	 _, err := cmd.Output()
	 check(err)	
}