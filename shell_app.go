package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)


/*
I hate doing this in this way.  It is suboptimal, I am planning to come back and do this with Map and without iteratios.
*/

func updateShellApp(input Input) {
	shellPath := REPOS_DIRS_MAP["SH"]
	branch := input.branch
	filePath := frontEndAppFullPath() + "/" + shellPath + "/" + "package.json"
	file := buildFileArray(filePath)
	for _, repo := range input.repos {
		if packageName, ok := REPOS_DIRS_MAP[repo]; ok {
			if ok {			
				updatedFile := findLine(file, packageName, branch)	
				output := strings.Join(updatedFile, "\n")								
				fmt.Println("Updated package json for " + repo)
				err := ioutil.WriteFile(filePath, []byte(output), 0644)
				check(err)			
			} else {
				fmt.Println(packageName + " was not found")
			}
		} else {
			fmt.Println(repo + " Not found.")
		}
	}	
}

func updatePackageJSON(path string, branch string, repo string) {

}

func buildFileArray(path string) []string {
	input, err := ioutil.ReadFile(path)
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
