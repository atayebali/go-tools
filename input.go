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
	//branchStr := flag.String("branch", "-1", "(Required) Name of the branch for the new feature")
	//reposStr := flag.String("repos", "SH", "(Optional) Pass in using CSV style")

	createFeatureCommand := flag.NewFlagSet("make-feature", flag.ExitOnError)
	branchStr := createFeatureCommand.String("branch", "-1", "(Required) Name of the branch for the new feature")
	reposStr := createFeatureCommand.String("repos", "SH", "(Optional) Pass in using CSV style")

	// flag.Parse()

	if os.Args[1] == "-h" {
		fmt.Printf(doc)
		os.Exit(0)
	}

	switch os.Args[1] {
	case "make-feature":
		createFeatureCommand.Parse(os.Args[2:])
	case "push-feature":
		commitFeatureStr
	}

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

var doc string = `
pr-spin
Create Branches and Generate PRs for mutliple repos. 

Syntax: 
SETUP: This will clone all the FE apps into ~/FRONT_END_APPS:
pr-spin -setup y

BUILD A FEATURE: This takes a branch name and keys to the repos you need initialized
pr-spin -make-feature -branch feature/HTW-1111-fix-table -repos TW,FB,CT,SH

PUSH A FEATURE
pr-spin -push-feature -branch feature/HTW-1111-fix-table -repos TW,FB,CT,SH	


Commands:
=========
	-setup:  
		Clones all the Front End Repos into ~/FRONT_END_APPS

	-make-feature:
		Creates a branch off master and pushes to remote. Updates package.json for Shell App
		makes a new yarn.lock and creates a PR.
		
		-branches:
			The branch that will be cut from master
			
		-repos:
			List of CSV keys corresponding the apps you need updating. 

	-push-feature:
		Commits changes on the branches you are working on and pushes to GH 
		creates	PR with master as base for all apps. 

		-branches:
			The branch that will be cut from master
			
		-repos:
			List of CSV keys corresponding the apps you need updating. 
`
