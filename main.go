package main

func main() {
	//Grab User Input with branch and repos
	input := parse()
	
	//Cut Branches off master
	cutBranch(input)

	// for _, repo := range input.repos {
	// 	fmt.Println(repo)
	// }
	// // lines, serr := readLines(".data")
	// check(serr)

	// dirPaths := filePaths(lines)
	// parallelLs(dirPaths)
}
