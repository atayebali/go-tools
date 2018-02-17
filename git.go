package main

import (
	"sync"
	"fmt"
)

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
}

func commitAndPush(input Input) {
	branch := input.branch
	repos := input.repos
	message := "Commit for " + branch

	paths := repoToPath(repos)
	projects := len(paths)
	var wg sync.WaitGroup
	wg.Add(projects)
	for _, path := range paths {
		go gitCommitWorker(&wg, message, path)
	}
	wg.Wait()

}

func generatePrs(input Input) {	
	fmt.Println(input)
	branch := input.branch
	repos := input.repos
	number := len(repos)
	var wg sync.WaitGroup
	wg.Add(number)
	for _, repo := range repos {
		githubName := REPOS_DIRS_MAP[repo]
		go gitHubWorker(&wg, RepoOpts{branch: branch, path: githubName})
	}
	wg.Wait()
}
