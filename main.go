package main

func main() {
	//Grab User Input with branch and repos
	input := parse()

	//Cut Branches off master
	cutBranchAndPush(input)
	
	//Prep Shell App
	updateShellApp(input)
}
