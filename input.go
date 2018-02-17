package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func parse() Input {
	//commandStr := flagString("")
	//prStr := flag.Bool("pr", false, "for testing purposes will run gh pr generation")

	input := Input{branch: "", repos: nil}
	createFeatureStr := flag.Bool("make-feature", false, "build a feature today!")
	commitFeatureStr := flag.Bool("push-feature", false, "push changes to github")
	setupStr := flag.String("setup", FRONTEND_APPS_BASE_DIR, "(Optional) Sets up the FRONT_END_APPS dir in your Home Dir")
	branchStr := flag.String("branch", "-1", "(Required) Name of the branch for the new feature")
	reposStr := flag.String("repos", "SH", "(Optional) Pass in using CSV style")
	flag.Parse()

	if *setupStr != FRONTEND_APPS_BASE_DIR {
		fmt.Println("Setting up Projects at default")
		setUp()
		os.Exit(0)
	}

	if *createFeatureStr {
		input = buildInput(input, *branchStr, *reposStr)
		input.command = "make-feature"
	}

	if *commitFeatureStr {
		input = buildInput(input, *branchStr, *reposStr)
		input.command = "push-feature"
	}

	return input
}

func buildInput(input Input, branch string, reposString string) Input {
	if branch == "-1" {
		fmt.Println("Need to pass branch name or else I fail!")
		os.Exit(1)
	}

	r := csv.NewReader(strings.NewReader(reposString))
	repos, err := r.Read()
	if err != nil {
		os.Exit(1)
	}
	input.branch = branch
	input.repos = repos
	return input
}
