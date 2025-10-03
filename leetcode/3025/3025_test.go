package leetcode_3025

import "testing"

func TestNumberOfPairs(t *testing.T) {
	tests := []struct {
		name     string
		points   [][]int
		expected int
	}{
		{
			name: "Example 1: No valid pairs",
			points: [][]int{
				{1, 1}, {2, 2}, {3, 3},
			},
			expected: 0,
		},
		{
			name: "Example 2: Multiple valid pairs",
			points: [][]int{
				{6, 2}, {4, 4}, {2, 6},
			},
			expected: 2,
		},
		{
			name: "Example 3: Multiple valid pairs with straight",
			points: [][]int{
				{3, 1}, {1, 3}, {1, 1},
			},
			expected: 2,
		},
		{
			name: "Example 4: Single point",
			points: [][]int{
				{0, 0},
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numberOfPairs(tt.points)
			if result != tt.expected {
				t.Errorf("numberOfPairs(%v) = %d, want %d", tt.points, result, tt.expected)
			}
		})
	}
}
