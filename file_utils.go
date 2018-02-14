package main

import (
	"bufio"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		os.Exit(1)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
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