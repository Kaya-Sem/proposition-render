package main

import (
	"testing"
)

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
		result := replaceAll(test.input, propositionSymbols)
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

func TestNormalizeParentheses(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"( )", "()"},
		{"( x -> y )", "(x -> y)"},
	}
	for _, test := range tests {
		result := normalizeParentheses(test.input)
		if result != test.expected {
			t.Errorf("normalizeWhitespace(%q) = %q; want %q", test.input, result, test.expected)
		}
	}

}
