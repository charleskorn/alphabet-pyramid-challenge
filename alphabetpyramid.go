package main

import (
	"os"
	"regexp"
)

type OutputWriter interface {
	WriteString(s string) (n int, err error)
}

func main() {
	run(os.Args[1:], os.Stdout)
}

func run(args []string, output OutputWriter) {
	if !validateArguments(args) {
		output.WriteString("INVALID INPUT\n")
		return
	}
}

func validateArguments(args []string) bool {
	if len(args) != 1 {
		return false
	}

	pattern := regexp.MustCompile(`^[A-Za-z]$`)

	return pattern.MatchString(args[0])
}
