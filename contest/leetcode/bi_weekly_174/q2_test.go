package biweekly174

import "testing"

func TestMinOperations(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3},
			target:   []int{2, 1, 3},
			expected: 2,
		},
		{
			name:     "Example 2",
			nums:     []int{4, 1, 4},
			target:   []int{5, 1, 4},
			expected: 1,
		},
		{
			name:     "Example 3",
			nums:     []int{7, 3, 7},
			target:   []int{5, 5, 9},
			expected: 2,
		},
		{
			name:     "Already equal",
			nums:     []int{1, 2, 3},
			target:   []int{1, 2, 3},
			expected: 0,
		},
		{
			name:     "All different same value",
			nums:     []int{5, 5, 5},
			target:   []int{1, 2, 3},
			expected: 1,
		},
		{
			name:     "All different different values",
			nums:     []int{1, 2, 3},
			target:   []int{4, 5, 6},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minOperations(tt.nums, tt.target)
			if result != tt.expected {
				t.Errorf("minOperations() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
