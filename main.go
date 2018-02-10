package main

import (
	"sync"
	"fmt"
)

func parallelLs(dirPaths []string) {
	sliceLength := len(dirPaths)
	var wg sync.WaitGroup
	wg.Add(sliceLength)
	for i := 0; i < sliceLength; i++ {
		go func(i int) {
			defer wg.Done()
			runGit(dirPaths, i)
		}(i)
	}
	wg.Wait()
}

func main() {
	input := parse()
	fmt.Println(input.branch)
	fmt.Println(input.repos)

	for _, repo := range input.repos {
		fmt.Println(repo)
	}
	// lines, serr := readLines(".data")
	// check(serr)
	// dirPaths := filePaths(lines)
	// parallelLs(dirPaths)
}
