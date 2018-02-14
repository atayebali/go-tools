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

func gitPushWorker(paths []string, branch string) {
	var wp sync.WaitGroup
	for _, path := range paths {
		wp.Add(1)
		go func(path string) {
			defer wp.Done()
			gitPush(RepoOpts{path: path, branch: branch})
		}(path)
	}
	wp.Wait()
}
