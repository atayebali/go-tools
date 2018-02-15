package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//SubOptimal brute force for now.  I have to find a good
//Lib to handle json and them we can make this work.

func updateShellApp(input Input) {
	shellPath := REPOS_DIRS_MAP["SH"]
	branch := input.branch
	repos := input.repos
	for i, repo := range repos {
		if string(repo[i]) != "SH" {
			updatePackageJson(shellPath, branch, repos[i])
		}
	}
}

func updatePackageJson(path string, branch string, repo string) {
	if packageName, ok := REPOS_DIRS_MAP[repo]; ok {
		if ok {
			file := buildFileArray(path)
			updatedFile := findLine(file, packageName, branch)
			output := strings.Join(updatedFile, "\n")
			err := ioutil.WriteFile("package.json", []byte(output), 0644)
			check(err)
		} else {
			fmt.Println(packageName + " was not found")
		}
	}
}

func buildFileArray(path string) []string {
	input, err := ioutil.ReadFile("package.json")
	check(err)
	rawFile := string(input)
	return strings.Split(rawFile, "\n")
}

func findLine(file []string, packageName string, branch string) []string {
	for i, line := range file {
		if strings.Contains(line, packageName) {
			result := strings.Split(line, "#")
			updatedLine := (result[0] + "#" + branch + "\",")
			file[i] = updatedLine
		}
	}
	return file
}
