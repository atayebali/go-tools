package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func splat(line string) {
	fmt.Println(line)
}

func filePaths(vars []string) []string {
	var paths []string
	for _, element := range vars {
		var tokens = strings.Split(element, "=")
		path := tokens[len(tokens)-1]
		paths = append(paths, path)
	}
	return paths
}

func parallelLs(dirPaths []string) {
	sliceLength := len(dirPaths)
	var wg sync.WaitGroup
	wg.Add(sliceLength)
	for i := 0; i < sliceLength; i++ {
		go func(i int) {
			defer wg.Done()
			
			cmd := exec.Command("git", "status")
			cmd1 := exec.Command("git", "rev-parse", "HEAD")
			cmd.Dir = dirPaths[i]
			cmd1.Dir = dirPaths[i]
			sha, err2 := cmd1.Output()
			out, err := cmd.Output()
			if err != nil {
				fmt.Println(err)
			}
			if err2 != nil {
				fmt.Println(err2)
			}


			fmt.Printf("The repo at %s is %s\n with sha @ %s", dirPaths[i], out, sha)
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
