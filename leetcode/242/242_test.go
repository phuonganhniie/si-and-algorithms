package leetcode_242

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{
			name:     "Example anagram",
			s:        "ab",
			t:        "a",
			expected: false,
		},
		{
			name:     "Example anagram",
			s:        "anagram",
			t:        "nagaram",
			expected: true,
		},
		{
			name:     "Example not anagram",
			s:        "rat",
			t:        "car",
			expected: false,
		},
		{
			name:     "Simple reorder",
			s:        "abc",
			t:        "cba",
			expected: true,
		},
		{
			name:     "Duplicates correct counts",
			s:        "aabb",
			t:        "bbaa",
			expected: true,
		},
		{
			name:     "Duplicates mismatch counts",
			s:        "aabb",
			t:        "abbb",
			expected: false,
		},
		{
			name:     "Different lengths",
			s:        "a",
			t:        "aa",
			expected: false,
		},
		{
			name:     "Single char same",
			s:        "z",
			t:        "z",
			expected: true,
		},
		{
			name:     "All same chars repeated",
			s:        "aaaa",
			t:        "aaaa",
			expected: true,
		},
		{
			name:     "Anagram with duplicates mixed",
			s:        "aab",
			t:        "aba",
			expected: true,
		},
		{
			name:     "Same length different composition",
			s:        "abc",
			t:        "abd",
			expected: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := isAnagram2(tt.s, tt.t)
			if got != tt.expected {
				t.Errorf("isAnagram(%q,%q) = %v, expected %v", tt.s, tt.t, got, tt.expected)
			}
		})
	}
}
