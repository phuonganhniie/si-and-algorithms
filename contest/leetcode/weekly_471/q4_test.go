package weekly471

import "testing"

func TestSumOfAncestorPairs(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		edges    [][]int
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			n:        3,
			edges:    [][]int{{0, 1}, {1, 2}},
			nums:     []int{2, 8, 2},
			expected: 3,
		},
		{
			name:     "Example 2",
			n:        3,
			edges:    [][]int{{0, 1}, {0, 2}},
			nums:     []int{1, 2, 4},
			expected: 1,
		},
		{
			name:     "Example 3",
			n:        4,
			edges:    [][]int{{0, 1}, {0, 2}, {1, 3}},
			nums:     []int{1, 2, 9, 4},
			expected: 2,
		},
		{
			name:     "Single child",
			n:        2,
			edges:    [][]int{{0, 1}},
			nums:     []int{4, 9},
			expected: 1, // 4 * 9 = 36 which is 6^2, so it IS a perfect square
		},
		{
			name:     "All same square-free",
			n:        4,
			edges:    [][]int{{0, 1}, {1, 2}, {2, 3}},
			nums:     []int{2, 2, 2, 2},
			expected: 6, // node 1: 1 ancestor, node 2: 2 ancestors, node 3: 3 ancestors = 1+2+3=6
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sumOfAncestorPairs(tt.n, tt.edges, tt.nums)
			if result != tt.expected {
				t.Errorf("sumOfAncestorPairs(%d, %v, %v) = %d, want %d",
					tt.n, tt.edges, tt.nums, result, tt.expected)
			}
		})
	}
}
