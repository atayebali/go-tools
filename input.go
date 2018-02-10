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
	branchStr := flag.String("branch", "", " (Required) Name of the branch for the new feature")
	reposStr := flag.String("repos", REPOS, "(Optional) Pass in using CSV style")
	flag.Parse()

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
