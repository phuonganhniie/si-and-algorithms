package leetcode_128

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Basic consecutive sequence",
			input:    []int{100, 4, 200, 1, 3, 2}, // sequence: 1,2,3,4 => length 4
			expected: 4,
		},
		{
			name:     "Single element array",
			input:    []int{1},
			expected: 1,
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "No consecutive elements",
			input:    []int{1, 3, 5, 7, 9}, // all isolated elements
			expected: 1,
		},
		{
			name:     "All same elements",
			input:    []int{5, 5, 5, 5}, // duplicates should be treated as one
			expected: 1,
		},
		{
			name:     "Two separate sequences",
			input:    []int{1, 2, 3, 10, 11, 12, 13}, // sequences: (1,2,3)=3, (10,11,12,13)=4
			expected: 4,
		},
		{
			name:     "Negative numbers sequence",
			input:    []int{-3, -2, -1, 0, 1}, // sequence: -3,-2,-1,0,1 => length 5
			expected: 5,
		},
		{
			name:     "Mixed positive and negative",
			input:    []int{-1, 1, 0, -2, 2}, // sequence: -2,-1,0,1,2 => length 5
			expected: 5,
		},
		{
			name:     "Large numbers with gaps",
			input:    []int{1000000, 1000001, 1000002, 500000}, // sequence: (1000000,1000001,1000002)=3
			expected: 3,
		},
		{
			name:     "Unordered with duplicates",
			input:    []int{9, 1, 4, 7, 3, 2, 8, 5, 6, 1, 9}, // sequence: 1,2,3,4,5,6,7,8,9 => length 9
			expected: 9,
		},
		{
			name:     "Reverse order consecutive",
			input:    []int{5, 4, 3, 2, 1}, // sequence: 1,2,3,4,5 => length 5
			expected: 5,
		},
		{
			name:     "Multiple small sequences",
			input:    []int{1, 2, 10, 11, 20, 21, 22}, // sequences: (1,2)=2, (10,11)=2, (20,21,22)=3
			expected: 3,
		},
		{
			name:     "Zero in the middle",
			input:    []int{-1, 0, 1}, // sequence: -1,0,1 => length 3
			expected: 3,
		},
		{
			name:     "Large array with one long sequence",
			input:    []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, // sequence: 0,1,2,3,4,5,6,7,8 => length 9
			expected: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestConsecutive(tt.input)
			if got != tt.expected {
				t.Errorf("longestConsecutive(%v) = %d, expected %d", tt.input, got, tt.expected)
			}
		})
	}
}
