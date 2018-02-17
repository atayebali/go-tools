package main

func main() {
	//Grab User Input with branch and repos
	input := parse()
	//Cut Branches off master
	if input.command == "make-feature" {
		cutBranchAndPush(input)

		//Prep Shell App
		updateShellApp(input)

		//Run Yarn Commands
		runYarnCommands()

		//Push Shell App
		gitPushShellApp(input)

		//Spin Up a PR
		opts := RepoOpts{branch: input.branch, path: REPOS_DIRS_MAP["SH"]}
		spinUpPr(opts)
	}

	if input.command == "push-feature" {
		commitAndPush(input)
		generatePrs(input)
	}
}
