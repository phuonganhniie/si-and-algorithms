package weekly_470

import "testing"

func TestLongestSubsequence(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name: "Remove 0",
			nums: []int{5, 3, 6, 0},
			expected: 3,
		},
		{
			name:     "Example 1: [1,2,3]",
			nums:     []int{1, 2, 3},
			expected: 2,
		},
		{
			name:     "Example 2: [2,3,4]",
			nums:     []int{2, 3, 4},
			expected: 3,
		},
		{
			name:     "All zeros",
			nums:     []int{0, 0, 0},
			expected: 0,
		},
		{
			name:     "Single non-zero element",
			nums:     []int{5},
			expected: 1,
		},
		{
			name:     "Single zero element",
			nums:     []int{0},
			expected: 0,
		},
		{
			name:     "XOR equals zero, remove one element",
			nums:     []int{1, 1},
			expected: 1,
		},
		{
			name:     "Mixed with zeros",
			nums:     []int{0, 1, 2, 0, 3},
			expected: 4,
		},
		{
			name:     "Large numbers",
			nums:     []int{1000000000, 999999999},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestSubsequence(tt.nums)
			if result != tt.expected {
				t.Errorf("longestSubsequence(%v) = %d; expected %d", tt.nums, result, tt.expected)
			}
		})
	}
}
