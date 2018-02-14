package main

import (
	"fmt"
	"os/exec"
	"sync"
	"gopkg.in/src-d/go-git.v4"
	. "gopkg.in/src-d/go-git.v4/_examples"
)

type RepoOpts struct {
	path   string
	branch string
}

/*
 args: Takes a dir path
 Moves contents into stash so that changes on a dirty branch are not lost
*/

func cutBranchAndPush(input Input) {
	branch := input.branch
	repos := input.repos

	paths := repoToPath(repos)
	projects := len(paths)
	var wg sync.WaitGroup	
	wg.Add(projects)
	for _, path := range paths {		
		go preBranchWorker(&wg, path)
	}
	wg.Wait()
	
	wg.Add(projects)
	for _, path := range paths {		
		go gitBranchWorker(&wg, RepoOpts{path: path, branch: branch} )
	}
	wg.Wait()

		//gitBranchWorker(paths, branch)
	// gitPushWorker(paths, branch)
}

func stashIt(opts RepoOpts) {
	cmd := exec.Command("git", "stash", "save", "-u")
	cmd.Dir = opts.path
	_, err := cmd.Output()
	CheckIfError(err)
	fmt.Printf("Stashing changes at %s \n", opts.path)
}

/*
args: path
Moves to branch, with Master as default
*/
func switchToBranch(opts RepoOpts) {
	r, err := git.PlainOpen(opts.path)
	CheckIfError(err)

	w, err := r.Worktree()
	CheckIfError(err)

	err1 := w.Checkout(&git.CheckoutOptions{})
	CheckIfError(err1)
	fmt.Println("Switching to master for " + opts.path)
}



/*
 args: path
 Pulls master from orgin
*/
func pullMaster(opts RepoOpts) {
	cmd := exec.Command("git", "pull", "origin", "master")
	cmd.Dir = opts.path
	_, err := cmd.Output()
	CheckIfError(err)
	fmt.Println("Pulled in Master: for " + opts.path)
}

func gitPush(opts RepoOpts) {
	cmd := exec.Command("git", "push", "origin", opts.branch)
	cmd.Dir = opts.path
	_, err := cmd.Output()
	CheckIfError(err)
	fmt.Println("Pushed to github: " + opts.branch)
}



func repoToPath(repoKeys []string) []string {
	var paths []string
	baseDir := getUserDir() + FRONTEND_APPS_PATH
	for _, repoKey := range repoKeys {
		if val, ok := REPOS_DIRS_MAP[repoKey]; ok {
			if ok {
				fullPath := baseDir + "/" + val
				paths = append(paths, fullPath)
			} else {
				fmt.Println("No Repo for " + repoKey)
			}
		}
	}
	return paths
}
