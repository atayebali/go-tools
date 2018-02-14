package main

import (
	"sync"
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

func gitBranchWorker(paths []string, branch string) {
	var wg sync.WaitGroup
	for _, path := range paths {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			gitBranch(RepoOpts{path: path, branch: branch})
		}(path)
	}
	wg.Wait()
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
