package leetcode_1304

import (
	"slices"
	"testing"
)

func TestSumZero(t *testing.T) {
	tests := []struct{
		name string
		n int
		expected []int
	}{
		{
			name: "Example 1",
			n: 5,
			expected: []int{1, -1, 2, -2, 0},
		},
		{
			name: "Example 2",
			n: 3,
			expected: []int{-1, 0, 1},
		},
		{
			name: "Example 3",
			n: 1,
			expected: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sumZero(tt.n)
			if slices.Equal(result, tt.expected) {
				t.Errorf("sumZero() = %d, want %d", result, tt.expected)
			}
		})
	}
}