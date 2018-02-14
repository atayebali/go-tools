package main

import (
	"sync"
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
		go gitBranchWorker(&wg, RepoOpts{path: path, branch: branch})
	}
	wg.Wait()

	wg.Add(projects)
	for _, path := range paths {
		go gitPushWorker(&wg, RepoOpts{path: path, branch: branch})
	}
	wg.Wait()
	// gitPushWorker(paths, branch)
}
