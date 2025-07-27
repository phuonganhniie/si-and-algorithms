package leetcode_3487

import "testing"

func TestMaxSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "All negative",
			nums:     []int{-1, -2, -3},
			expected: -1,
		},
		{
			name:     "Positive and negative",
			nums:     []int{-20, 20},
			expected: 20,
		},
		{
			name:     "Example 1: All unique elements",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "Example 2: Many duplicates",
			nums:     []int{1, 1, 0, 1, 1},
			expected: 1,
		},
		{
			name:     "Example 3: Mixed with negatives",
			nums:     []int{1, 2, -1, -2, 1, 0, -1},
			expected: 3,
		},
		{
			name:     "Single element",
			nums:     []int{5},
			expected: 5,
		},
		{
			name:     "Zeros and positives",
			nums:     []int{0, 1, 0, 2, 0},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxSum(tt.nums)
			if result != tt.expected {
				t.Errorf("maxSum(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
