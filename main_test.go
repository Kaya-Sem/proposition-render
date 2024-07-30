package main

import (
	"fmt"
	"testing"
)

var clipboardWriteAll = func(s string) error {
	return nil
}

// Test replaceAll function
func TestReplaceAll(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"a && b", "a ∧ b"},
		{"a || b", "a ∨ b"},
		{"a == b", "a ≡ b"},
		{"a != b", "a ≢ b"},
		{"a -> b", "a ➟ b"},
		{"forall x", "∀ x"},
		{"exists y", "∃ y"},
		{"a and b", "a ⩓ b"},
		{"a && b || c", "a ∧ b ∨ c"},
	}

	for _, test := range tests {
		result := replaceAll(test.input)
		if result != test.expected {
			t.Errorf("replaceAll(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

// Test normalizeWhitespace function
func TestNormalizeWhitespace(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"a    b", "a b"},
		{"a  b   c", "a b c"},
		{"   a   b   c   ", "a b c"},
		{"a\nb\nc", "a b c"},
		{"a\tb\tc", "a b c"},
	}

	for _, test := range tests {
		result := normalizeWhitespace(test.input)
		if result != test.expected {
			t.Errorf("normalizeWhitespace(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

//Mocking the clipboard functionality
// You can use the following code if you want to test copyToClipboard function

// Test copyToClipboard function
func TestCopyToClipboard(t *testing.T) {
	// Successful copy
	clipboardWriteAll = func(s string) error {
		if s == "expected text" {
			return nil
		}
		return fmt.Errorf("unexpected text")
	}

	err := copyToClipboard("expected text")
	if err != nil {
		t.Errorf("copyToClipboard() = %v; want no error", err)
	}

	// Failure case
	clipboardWriteAll = func(s string) error {
		return fmt.Errorf("clipboard error")
	}

	err = copyToClipboard("unexpected text")
	if err == nil {
		t.Errorf("copyToClipboard() = no error; want an error")
	}
}
