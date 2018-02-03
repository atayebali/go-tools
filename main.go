package main

import (
	"sync"
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
	
	lines, serr := readLines(".data")
	check(serr)
	dirPaths := filePaths(lines)
	parallelLs(dirPaths)
}
