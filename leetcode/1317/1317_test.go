package leetcode_1317

import (
	"slices"
	"testing"
)

func TestGetNoZeroIntegers(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected []int
	}{
		{
			name:     "Example 1",
			n:        2,
			expected: []int{1, 1},
		},
		{
			name:     "Example 2",
			n:        11,
			expected: []int{2, 9},
		},
		{
			name:     "Large number",
			n:        101,
			expected: []int{2, 99},
		},
		{
			name:     "Edge case 1",
			n:        3,
			expected: []int{1, 2},
		},
		{
			name:     "Edge case 2",
			n:        1001,
			expected: []int{2, 999},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getNoZeroIntegers(tt.n)
			if !slices.Equal(result, tt.expected) {
				t.Errorf("getNoZeroIntegers() = %v, want %v", result, tt.expected)
			}
		})
	}
}
