package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	utilities "pcheck/util"
	"regexp"
	"strings"
)

func scanWords(reader io.Reader) []string {
	words := []string{}

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		utilities.CheckIfError(err)
	}

	return words
}

func isURL(path string) bool {
	return strings.HasPrefix(path, "http")
}

func readFromURL(url string) []string {
	resp, err := http.Get(url)
	utilities.CheckIfError(err)
	defer resp.Body.Close()

	rawData, err := ioutil.ReadAll(resp.Body)
	utilities.CheckIfError(err)

	if json.Valid(rawData) {

		var words []string
		utilities.CheckIfError(json.Unmarshal(rawData, &words))
		return words
	}

	// Assume it's a list of words
	return scanWords(resp.Body)
}

func importWordlist(pathToWordlist string) []string {

	if isURL(pathToWordlist) {
		return readFromURL(pathToWordlist)
	}

	list, err := os.Open(pathToWordlist)
	utilities.CheckIfError(err)
	defer list.Close()

	return scanWords(list)
}

func wordExists(needle string, haystack string) bool {

	obj, err := regexp.Match(regexp.QuoteMeta(needle), bytes.NewBufferString(haystack).Bytes())
	utilities.CheckIfError(err)

	return obj
}
