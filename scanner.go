package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	utilities "pcheck/util"
	"strings"

	"github.com/go-git/go-git"
	"github.com/go-git/go-git/plumbing/object"
)

func checkGitCommits(gitPath string, wordList []string) {
	// We instantiate a new repository targeting the given path (the .git folder)
	r, err := git.PlainOpen(gitPath)
	utilities.CheckIfError(err)

	// ... retrieving the HEAD reference
	ref, err := r.Head()
	utilities.CheckIfError(err)

	// ... retrieves the commit history
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
	utilities.CheckIfError(err)

	// ... just iterates over the commits
	var wordCounter int
	err = cIter.ForEach(func(c *object.Commit) error {

		for _, word := range wordList {
			if wordExists(word, c.String()) {
				fmt.Println("Found ", word)
				wordCounter++
			}
		}
		return nil
	})
	utilities.CheckIfError(err)

	printResults(wordCounter, "git logs")

}

func checkFilesAtPath(gitPath string, wordList []string) {
	err := filepath.Walk(gitPath,
		func(path string, info os.FileInfo, err error) error {
			var tempCounter int

			if !info.IsDir() && !strings.Contains(path, ".git") {
				utilities.CheckIfError(err)
				data := checkFile(path)

				for _, word := range wordList {
					if wordExists(word, data) {
						fmt.Println("Found ", word)
						tempCounter++
					}
				}

				printResults(tempCounter, path)
			}
			return nil
		})
	utilities.CheckIfError(err)
}

func checkFile(path string) string {

	file, err := os.Open(path)
	utilities.CheckIfError(err)
	defer file.Close()

	rawData, err := ioutil.ReadAll(file)
	utilities.CheckIfError(err)

	return string(rawData)

}

func printResults(count int, path string) {
	if count > 0 {
		fmt.Println(count, "found in", path)
	}
}
