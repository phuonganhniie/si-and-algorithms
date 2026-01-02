package leetcode_961

import (
	"fmt"
	"testing"
)

func TestRepeatedNTimes(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "Example 1: [1,2,3,3]",
			input: []int{1, 2, 3, 3},
			want:  3,
		},
		{
			name:  "Example 2: [2,1,2,5,3,2]",
			input: []int{2, 1, 2, 5, 3, 2},
			want:  2,
		},
		{
			name:  "Example 3: [5,1,5,2,5,3,5,4]",
			input: []int{5, 1, 5, 2, 5, 3, 5, 4},
			want:  5,
		},
		{
			name:  "Minimum size: n=1",
			input: []int{1, 1},
			want:  1,
		},
		{
			name:  "Repeated at start",
			input: []int{9, 9, 9, 1},
			want:  9,
		},
		{
			name:  "Repeated at end",
			input: []int{1, 7, 7, 7},
			want:  7,
		},
		{
			name:  "Repeated element is 0",
			input: []int{0, 0, 1, 2},
			want:  0,
		},
		{
			name:  "Large repeated value",
			input: []int{1000, 2000, 3000, 1000},
			want:  1000,
		},
		{
			name:  "Scattered repetitions",
			input: []int{4, 1, 4, 2, 4, 3},
			want:  4,
		},
		{
			name:  "All same except one",
			input: []int{8, 8, 8, 8, 8, 1},
			want:  8,
		},
		{
			name:  "Negative numbers",
			input: []int{-1, -1, -1, 5},
			want:  -1,
		},
		{
			name:  "Mixed positive and negative",
			input: []int{-5, 3, -5, 7, -5, 9},
			want:  -5,
		},
		{
			name:  "Adjacent repetitions",
			input: []int{6, 6, 6, 10, 11, 12},
			want:  6,
		},
		{
			name:  "Non-adjacent repetitions",
			input: []int{1, 3, 5, 3, 7, 3},
			want:  3,
		},
		{
			name:  "Larger array",
			input: []int{2, 6, 2, 8, 2, 10, 2, 12},
			want:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := repeatedNTimes(tt.input)
			if got != tt.want {
				t.Errorf("repeatedNTimes(%v) = %d, want %d", tt.input, got, tt.want)
				fmt.Printf("Input: %v\n", tt.input)
			} else {
				fmt.Printf("✓ %s: PASSED (input=%v, output=%d)\n", tt.name, tt.input, got)
			}
		})
	}
}
