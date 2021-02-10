package utilities

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/akamensky/argparse"
)

// CheckArgs sets arguments for the application.
func CheckArgs() map[string]*string {

	args := make(map[string]*string)

	parser := argparse.NewParser("print", "Prints provided string to stdout")

	args["git_path"] = parser.String("p", "git_path", &argparse.Options{Required: true, Help: "Path to git repository"})
	args["word_list"] = parser.String("w", "word_list", &argparse.Options{Required: true, Help: "Path to word list. HTTPS/HTTP/FS"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	return args
}

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	debug.PrintStack()
	os.Exit(1)
}

// Info should be used to describe the example commands that are about to run.
func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// Warning should be used to display a warning
func Warning(format string, args ...interface{}) {
	fmt.Printf("\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}
