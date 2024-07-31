package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"os"
	"strings"
)

// symbols maps standard logical operators and quantifiers to their Unicode equivalents.
var symbols = map[string]string{
	"&&":     "∧",
	"||":     "∨",
	"==":     "≡",
	"!=":     "≢",
	"->":     "➟",
	"forall": "∀",
	"exists": "∃",
	"and":    "⩓",
}

const NEEDED_ARGS int = 2

func main() {
	if len(os.Args) < NEEDED_ARGS {
		fmt.Println("Usage: proposition \"<proposition>\"")
		return
	}

	proposition := strings.Join(os.Args[1:], " ")
	replaced := replaceAll(proposition)
	replaced = normalizeWhitespace(replaced)
	replaced = normalizeParentheses(replaced)
	fmt.Println(replaced)

	copyToClipboard(replaced)

}

func copyToClipboard(input string) {
	err := clipboard.WriteAll(input)
	if err != nil {
		fmt.Println("Error copying to clipboard:", err)
	} else {
		fmt.Println("\nProposition has been copied to clipboard")
	}
}

// replaceAll replaces logical operators and quantifiers with their Unicode equivalents.
func replaceAll(proposition string) string {
	for key, value := range symbols {
		proposition = strings.Replace(proposition, key, value, -1)
	}

	return proposition
}

// TODO: create tests for this function.

const OCCURANCES = -1 // -1 is all occurances

func normalizeParentheses(s string) string {
	s = strings.Replace(s, "( ", "(", OCCURANCES)
	s = strings.Replace(s, " )", ")", OCCURANCES)
	return s
}

// replaces multiple consecutive whitespace  (and newline )characters with a single space.
const JOINING_STRING = " " // space

func normalizeWhitespace(s string) string {
	fields := strings.Fields(s)
	return strings.Join(fields, JOINING_STRING)
}
