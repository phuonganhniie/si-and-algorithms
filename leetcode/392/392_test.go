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

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := isSubsequence(tt.s, tt.t)
			if got != tt.expected {
				t.Errorf("isSubsequence(%q, %q) = %v, expected %v", tt.s, tt.t, got, tt.expected)
			}
		})
	}
}
