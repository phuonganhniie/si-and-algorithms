package leetcode_392

import "testing"

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{
			name:     "Valid subsequence - basic case",
			s:        "abc",
			t:        "ahbgdc",
			expected: true,
		},
		{
			name:     "Invalid subsequence - character not found in order",
			s:        "axc",
			t:        "ahbgdc",
			expected: false,
		},
		{
			name:     "Valid subsequence - consecutive characters",
			s:        "ace",
			t:        "abcde",
			expected: true,
		},
		{
			name:     "Invalid subsequence - wrong order",
			s:        "aec",
			t:        "abcde",
			expected: false,
		},
		{
			name:     "Empty subsequence",
			s:        "",
			t:        "abc",
			expected: true,
		},
		{
			name:     "Empty target string with non-empty subsequence",
			s:        "a",
			t:        "",
			expected: false,
		},
		{
			name:     "Both strings empty",
			s:        "",
			t:        "",
			expected: true,
		},
		{
			name:     "Single character match",
			s:        "a",
			t:        "a",
			expected: true,
		},
		{
			name:     "Subsequence longer than target",
			s:        "abc",
			t:        "ab",
			expected: false,
		},
	}

	// Test all approaches
	approaches := []struct {
		name string
		fn   func(string, string) bool
	}{
		{"Two Pointer", isSubsequence},
		{"Binary Search", isSubsequenceBinarySearch},
		{"Recursive", isSubsequenceRecursive},
		{"Bit Mask", isSubsequenceBitMask},
	}

	for _, approach := range approaches {
		approach := approach
		t.Run(approach.name, func(t *testing.T) {
			for _, tt := range tests {
				tt := tt
				t.Run(tt.name, func(t *testing.T) {
					got := approach.fn(tt.s, tt.t)
					if got != tt.expected {
						t.Errorf("%s: isSubsequence(%q, %q) = %v, expected %v",
							approach.name, tt.s, tt.t, got, tt.expected)
					}
				})
			}
		})
	}
}
