package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	"strings"
)

/*
I hate doing this in this way.  It is suboptimal, I am planning to come back and do this with Map and without iteratios.
*/
var shellDir string = REPOS_DIRS_MAP["SH"]
var shellPath string = frontEndAppFullPath() + "/" + shellDir
var filePath string = shellPath + "/" + "package.json"

func updateShellApp(input Input) {
	branch := input.branch
	file := buildFileArray(filePath)
	for _, repo := range input.repos {
		if repo == "SH" {
			continue
		}
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

func runYarnCommands() {
	fmt.Println("Killing node modeules")
	delNodeModules := exec.Command("rm", "-rf", "node_modules")
	delNodeModules.Dir = shellPath
	_, err := delNodeModules.Output()
	check(err)
	fmt.Println("Running yarn install")
	yarnInstall := exec.Command("yarn", "install")
	yarnInstall.Dir = shellPath
	stdout, err1 := yarnInstall.StdoutPipe()
	check(err1)
	yarnInstall.Start()
	oneByte := make([]byte, 100)
	for {
		_, err := stdout.Read(oneByte)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err.Error())
		}

		r := bufio.NewReader(stdout)
		line, _, _ := r.ReadLine()
		fmt.Println(string(line))
	}
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

func gitPushShellApp(input Input) {
	branch := input.branch
	gitAdd := exec.Command("git", "add", ".")
	gitAdd.Dir = shellPath

	_, err := gitAdd.Output()
	check(err)

	message := "Commit for feature " + branch
	gitCommit := exec.Command("git", "commit", "-am", message)
	gitCommit.Dir = shellPath

	_, err1 := gitCommit.Output()
	check(err1)

	opts := RepoOpts{path: shellPath, branch: branch}
	gitPush(opts)
}
