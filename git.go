package main

import (
	"fmt"
	"os/exec"
	"gopkg.in/src-d/go-git.v4"
	. "gopkg.in/src-d/go-git.v4/_examples"
)
type RepoOpts struct{
	path string
	branch string
}

/*
 args: Takes a dir path
 Moves contents into stash so that changes on a dirty branch are not lost
*/

func stashIt(opts RepoOpts){
	cmd := exec.Command("git", "stash")
	cmd.Dir = opts.path
	out, err := cmd.Output()
	CheckIfError(err)
	fmt.Printf("The stash is %s", out)
}

/*
args: path 
Moves to branch, with Master as default
*/
func switchToBranch(opts RepoOpts){
	r, err := git.PlainOpen(opts.path)
	CheckIfError(err)
	
	w, err := r.Worktree()
	CheckIfError(err)

	err1 := w.Checkout(&git.CheckoutOptions{})
	CheckIfError(err1)
}

/*
 args: path 
 Pulls master from orgin
*/
func pullMaster(opts RepoOpts){

}


func runGit(dirPaths []string, i int) {
	stashIt(RepoOpts{path: dirPaths[i]})
	switchToBranch(RepoOpts{path: dirPaths[i], branch: "Master"})

}
