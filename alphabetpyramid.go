package main

import (
	"os"
	"regexp"
	"strings"
	"unicode"
)

type OutputWriter interface {
	WriteString(s string) (n int, err error)
}

func main() {
	run(os.Args[1:], os.Stdout)
}

func run(args []string, output OutputWriter) {
	letter, valid := parseArguments(args)

	if !valid {
		output.WriteString("INVALID INPUT\n")
		return
	}

	output.WriteString(generatePyramid(letter))
}

func parseArguments(args []string) (rune, bool) {
	if len(args) != 1 {
		return 0, false
	}

	firstArgument := args[0]

	pattern := regexp.MustCompile(`^[A-Za-z]$`)

	if !pattern.MatchString(firstArgument) {
		return 0, false
	}

	letter := []rune(firstArgument)[0]

	return unicode.ToUpper(letter), true
}

func generatePyramid(lastLetter rune) string {
	totalLetters := int(lastLetter-'A') + 1
	pyramid := ""

	for currentLetter := 'A'; currentLetter < lastLetter; currentLetter++ {
		pyramid += generatePyramidLine(currentLetter, totalLetters) + "\n"
	}

	pyramid += generatePyramidLine(lastLetter, totalLetters) + "\n"

	for currentLetter := (lastLetter - 1); currentLetter >= 'A'; currentLetter-- {
		pyramid += generatePyramidLine(currentLetter, totalLetters) + "\n"
	}

	return pyramid
}

func generatePyramidLine(letter rune, totalLetters int) string {
	index := int(letter - 'A')
	outsideSpaces := totalLetters - index - 1

	if letter == 'A' {
		return strings.Repeat(" ", outsideSpaces) + string(letter)
	} else {
		insideSpaces := 2*index - 1

		return strings.Repeat(" ", outsideSpaces) + string(letter) + strings.Repeat(" ", insideSpaces) + string(letter)
	}

}
