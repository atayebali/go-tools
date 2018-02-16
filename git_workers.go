package main

import (
	"sync"
	"fmt"
	"os/exec"
)

func preBranchWorker(wg *sync.WaitGroup, path string) {
	defer wg.Done()
	preBranchingSteps(path)
}

func preBranchingSteps(path string) {
	opts := RepoOpts{path: path}
	stashIt(opts)
	switchToBranch(opts)
	pullMaster(opts)
}

func gitBranchWorker(wg *sync.WaitGroup, opts RepoOpts) {
	defer wg.Done()	
	gitBranch(opts)
}

func gitBranch(opts RepoOpts){
	cmd := exec.Command("git", "checkout", "-b", opts.branch)
	cmd.Dir = opts.path
	_, err := cmd.Output()

	if (err != nil) && (string(err.Error()) == "exit status 128") {
		cmd := exec.Command("git", "checkout", opts.branch)
		cmd.Dir = opts.path
		_, err := cmd.Output()
		check(err)
	}
	fmt.Println("New branch created in " + opts.path + "for " + opts.branch)
}

func gitPushWorker(wg *sync.WaitGroup, opts RepoOpts) {
	defer wg.Done()
	gitPush(opts)
}

func gitPush(opts RepoOpts) {
	cmd := exec.Command("git", "push", "origin", opts.branch)
	cmd.Dir = opts.path
	fmt.Println("Pushing to github: " + opts.branch)
	 _, err := cmd.Output()
	 check(err)
	
}
