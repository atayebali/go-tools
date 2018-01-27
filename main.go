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

	fmt.Println("Running for Ls...")

	for i := 0; i < sliceLength; i++ {
		go func(i int) {
			defer wg.Done()
			out, err := exec.Command("ls", "-la", dirPaths[i]).Output()
			if err != nil {
				panic(err)
			}
			fmt.Printf("The dir is %s\n", out)
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
