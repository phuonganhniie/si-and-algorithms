package weekly471

import "testing"

func TestLongestBalancedSubstring(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "Example 1",
			s:        "abbac",
			expected: 4,
		},
		{
			name:     "Example 2",
			s:        "aabcc",
			expected: 3,
		},
		{
			name:     "Example 3",
			s:        "aba",
			expected: 2,
		},
		{
			name:     "Single character",
			s:        "a",
			expected: 1,
		},
		{
			name:     "All same characters",
			s:        "aaaa",
			expected: 4,
		},
		{
			name:     "Two different characters balanced",
			s:        "aabb",
			expected: 4,
		},
		{
			name:     "Three different characters balanced",
			s:        "aabbcc",
			expected: 6,
		},
		{
			name:     "Complex case",
			s:        "abcabc",
			expected: 6,
		},
		{
			name:     "Unbalanced at start",
			s:        "aaabbc",
			expected: 4,
		},
		{
			name:     "Mixed balanced substrings",
			s:        "abccba",
			expected: 6,
		},
		{
			name:     "Test",
			s:        "accc",
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestBalancedSubstring(tt.s)
			if result != tt.expected {
				t.Errorf("longestBalancedSubstring(%q) = %d, want %d", tt.s, result, tt.expected)
			}
		})
	}
}
