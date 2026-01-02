package leetcode_1392

import (
	"fmt"
	"testing"
)

func TestLongestPrefixBruteForce(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Example 1: level",
			input: "level",
			want:  "l",
		},
		{
			name:  "Example 2: ababab",
			input: "ababab",
			want:  "abab",
		},
		{
			name:  "No happy prefix",
			input: "abc",
			want:  "",
		},
		{
			name:  "Single character",
			input: "a",
			want:  "",
		},
		{
			name:  "Two same characters",
			input: "aa",
			want:  "a",
		},
		{
			name:  "All same characters",
			input: "aaaa",
			want:  "aaa",
		},
		{
			name:  "Pattern at start and end",
			input: "ababcabab",
			want:  "abab",
		},
		{
			name:  "Short repeating pattern",
			input: "abab",
			want:  "ab",
		},
		{
			name:  "Long string with no match",
			input: "abcdefghij",
			want:  "",
		},
		{
			name:  "Partial overlap",
			input: "acccbaaacccb",
			want:  "acccb",
		},
		{
			name:  "Complex pattern 1",
			input: "aabaabaaa",
			want:  "aa",
		},
		{
			name:  "Complex pattern 2",
			input: "ababcabaab",
			want:  "ab",
		},
		{
			name:  "Three character pattern",
			input: "abcabc",
			want:  "abc",
		},
		{
			name:  "Nested pattern",
			input: "aabaaab",
			want:  "aab",
		},
		{
			name:  "Long repeating",
			input: "abcabcabc",
			want:  "abcabc",
		},
		{
			name:  "Single repeat at end",
			input: "abca",
			want:  "a",
		},
		{
			name:  "Multiple possible matches",
			input: "abacabad",
			want:  "",
		},
		{
			name:  "Edge case: two characters different",
			input: "ab",
			want:  "",
		},
		{
			name:  "Almost full string",
			input: "aaa",
			want:  "aa",
		},
		{
			name:  "Pattern with different lengths",
			input: "aaacaaaa",
			want:  "aaa",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestPrefix(tt.input)
			if got != tt.want {
				t.Errorf("longestPrefixBruteForce(%q) = %q, want %q", tt.input, got, tt.want)
				fmt.Printf("Input: %q\n", tt.input)
			} else {
				fmt.Printf("✓ %s: PASSED (input=%q, output=%q)\n", tt.name, tt.input, got)
			}
		})
	}
}
