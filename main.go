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

func main() {
	if len(os.Args) < 2 {
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
func normalizeParentheses(s string) string {
	s = strings.Replace(s, "( ", "(", -1)
	s = strings.Replace(s, " )", ")", -1)
	return s
}

// replaces multiple consecutive whitespace  (and newline )characters with a single space.
func normalizeWhitespace(s string) string {
	fields := strings.Fields(s)
	return strings.Join(fields, " ")
}
