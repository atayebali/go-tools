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
	packStr := flag.Bool("pack", false, "for testing purposes will run update packagejson")
	setupStr := flag.String("setup", FRONTEND_APPS_BASE_DIR, "(Optional) Sets up the FRONT_END_APPS dir in your Home Dir")
	branchStr := flag.String("branch", "-1", "(Required) Name of the branch for the new feature")
	reposStr := flag.String("repos", "SH", "(Optional) Pass in using CSV style")
	flag.Parse()

	if *packStr {
		runYarnCommands()
		os.Exit(0)
	}

	if *setupStr != FRONTEND_APPS_BASE_DIR {
		fmt.Println("Setting up Projects at default")
		setUp()
		os.Exit(0)
	}

	if *branchStr == "-1" {
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
