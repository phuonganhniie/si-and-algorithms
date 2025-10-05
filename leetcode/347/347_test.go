package leetcode_347

import (
	"reflect"
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "Oreilly case",
			nums:     []int{5, 5, 5, 6, 6, 7, 8, 9},
			k:        2,
			expected: []int{5, 6},
		},
		{
			name:     "Single element with k=1",
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
		{
			name:     "Multiple elements same frequency",
			nums:     []int{1, 2, 3, 4},
			k:        2,
			expected: []int{1, 2}, // or any 2 elements since all have same frequency
		},
		{
			name:     "Basic case with k=2",
			nums:     []int{1, 1, 1, 2, 2, 3},
			k:        2,
			expected: []int{1, 2},
		},
		{
			name:     "All same elements",
			nums:     []int{4, 4, 4, 4},
			k:        1,
			expected: []int{4},
		},
		{
			name:     "Negative numbers",
			nums:     []int{-1, -1, -1, 2, 2, 3},
			k:        2,
			expected: []int{-1, 2},
		},
		{
			name:     "Large frequency difference",
			nums:     []int{1, 1, 1, 1, 1, 2, 2, 3},
			k:        2,
			expected: []int{1, 2},
		},
		{
			name:     "k equals array length",
			nums:     []int{5, 6, 7},
			k:        3,
			expected: []int{5, 6, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(tt.nums, tt.k)

			// Sort both slices for comparison since order doesn't matter for same frequency
			sort.Ints(got)
			sort.Ints(tt.expected)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("topKFrequent(%v, %d) = %v, expected %v", tt.nums, tt.k, got, tt.expected)
			}
		})
	}
}
