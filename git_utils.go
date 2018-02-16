package main

import (
	"fmt"
	"os/exec"
	"gopkg.in/src-d/go-git.v4"
	. "gopkg.in/src-d/go-git.v4/_examples"
)

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


func repoToPath(repoKeys []string) []string {
	var paths []string
	baseDir := getUserDir() + FRONTEND_APPS_BASE_DIR
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
