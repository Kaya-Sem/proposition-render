package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"os"
	"strings"
)

// propositionSymbols maps standard logical operators and quantifiers to their Unicode equivalents.
var propositionSymbols = map[string]string{
	"&&":     "∧",
	"||":     "∨",
	"==":     "≡",
	"!=":     "≢",
	"->":     "➟",
	"forall": "∀",
	"exists": "∃",
	"and":    "⩓",
}

var latexSymbols = map[string]string{
	"&&":     "\\wedge",
	"||":     "\\vee",
	"==":     "\\eq",
	"!=":     "\\neq",
	"->":     "\\Rightarrow",
	"forall": "\\forall",
	"exists": "\\exists",
}

const NEEDED_ARGS int = 2
const REPLACEMENT_OCCURANCES = -1 // -1 is all occurances

func main() {
	if len(os.Args) < NEEDED_ARGS {
		fmt.Println("Usage: proposition \"<proposition>\"")
		return
	}

	proposition := strings.Join(os.Args[1:], " ")
	replaced := replaceAll(proposition, propositionSymbols)
	replaced = normalizeWhitespace(replaced)
	replaced = normalizeParentheses(replaced)

	latexProposition := proposition
	latexProposition = createLatexString(latexProposition)
	fmt.Println(latexProposition)

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
func replaceAll(proposition string, symbols_map map[string]string) string {
	for key, value := range symbols_map {
		proposition = strings.Replace(proposition, key, value, REPLACEMENT_OCCURANCES)
	}

	return proposition
}

// TODO: create tests for this function.

func normalizeParentheses(s string) string {
	s = strings.Replace(s, "( ", "(", REPLACEMENT_OCCURANCES)
	s = strings.Replace(s, " )", ")", REPLACEMENT_OCCURANCES)
	return s
}

// replaces multiple consecutive whitespace  (and newline )characters with a single space.
const JOINING_STRING = " " // space

func normalizeWhitespace(s string) string {
	fields := strings.Fields(s)
	return strings.Join(fields, JOINING_STRING)
}

func createLatexString(input string) string {
	input = normalizeWhitespace(input)
	input = normalizeParentheses(input)
	input = "$" + input + "$"
	input = replaceAll(input, latexSymbols)
	return input
}
