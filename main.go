package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"os"
	"strings"
)

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
		fmt.Println("Please provide a proposition string as an argument.")
		return
	}

	proposition := os.Args[1]
	replaced := replaceAll(proposition)
	replaced = normalizeWhitespace(replaced)
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

func replaceAll(proposition string) string {
	for key, value := range symbols {
		proposition = strings.Replace(proposition, key, value, -1)
	}

	return proposition
}

func normalizeWhitespace(s string) string {
	fields := strings.Fields(s)
	return strings.Join(fields, " ")
}
