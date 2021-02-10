package main

import (
	utilities "pcheck/util"
)

func main() {

	args := utilities.CheckArgs()

	wordlist := importWordlist(*args["word_list"])

	checkGitCommits(*args["git_path"], wordlist)
	checkFilesAtPath(*args["git_path"], wordlist)
}
