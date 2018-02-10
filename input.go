package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Input struct {
	branch string
	repos  []string
}

const REPOS = "PT,SC,TW,FB,SG,CT"

func parse() Input {
	setupStr := flag.String("setup", "mkdir", "Sets up the FRONT_END_APPS dir in your Home Dir")
	branchStr := flag.String("branch", "", " (Required) Name of the branch for the new feature")
	reposStr := flag.String("repos", REPOS, "(Optional) Pass in using CSV style")
	flag.Parse()

	if *setupStr != "" {
		fmt.Println("Setting up Projects")
		setUp()
		os.Exit(0)
	}

	if *branchStr == "" {
		fmt.Println("Need to pass branch name or else I fail!")
		os.Exit(1)
	}

	r := csv.NewReader(strings.NewReader(*reposStr))
	repos, err := r.Read()
	if err != nil {
		os.Exit(1)
	}
	input := Input{branch: *branchStr, repos: repos}
	return input

}
