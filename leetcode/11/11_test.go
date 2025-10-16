package leetcode_11

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "Zigzag pattern",
			height: []int{1, 5, 2, 5, 1, 5, 2, 5, 1},
			want:   30, // Area between index 1 and 7: 6 * 5 = 30, or others
		},
		{
			name:   "Descending order",
			height: []int{5, 4, 3, 2, 1},
			want:   6, // Area between index 0 (height=5) and index 4 (height=1): 4 * 1 = 4, or index 0 and 3: 3 * 2 = 6
		},
		{
			name:   "Large numbers",
			height: []int{100, 200, 150, 300, 250},
			want:   600, // Area between index 1 and 4: 3 * 200 = 600
		},
		{
			name:   "Example 1: Standard case with multiple heights",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49, // Area between index 1 (height=8) and index 8 (height=7): 7 * 7 = 49
		},
		{
			name:   "Example 2: Minimum case with two elements",
			height: []int{1, 1},
			want:   1, // Area: 1 * 1 = 1
		},
		{
			name:   "All same heights",
			height: []int{5, 5, 5, 5, 5},
			want:   20, // Area between first and last: 4 * 5 = 20
		},
		{
			name:   "Ascending order",
			height: []int{1, 2, 3, 4, 5},
			want:   6, // Area between index 0 (height=1) and index 4 (height=5): 4 * 1 = 4, or index 1 and 4: 3 * 2 = 6
		},
		{
			name:   "High walls at ends",
			height: []int{10, 1, 1, 1, 1, 1, 1, 1, 1, 10},
			want:   90, // Area between first and last: 9 * 10 = 90
		},
		{
			name:   "High walls in middle",
			height: []int{1, 10, 10, 1},
			want:   10, // Area between index 0 and 3: 3 * 1 = 3, or index 1 and 2: 1 * 10 = 10
		},
		{
			name:   "Peak in the middle",
			height: []int{1, 3, 2, 5, 25, 24, 5},
			want:   24, // Area between index 4 and 5: 1 * 24 = 24
		},
		{
			name:   "Multiple equal maximum heights",
			height: []int{8, 2, 3, 4, 8, 2, 8},
			want:   48, // Area between index 0 and 6: 6 * 8 = 48
		},
		{
			name:   "Two high walls close together",
			height: []int{1, 2, 1, 100, 100, 1, 2, 1},
			want:   100, // Area between index 3 and 4: 1 * 100 = 100
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxArea(tt.height)
			if got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
				fmt.Printf("Input: %v\n", tt.height)
			} else {
				fmt.Printf("✓ %s: PASSED (result: %d)\n", tt.name, got)
			}
		})
	}
}

// Benchmark test to measure performance
func BenchmarkMaxArea(b *testing.B) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	for i := 0; i < b.N; i++ {
		maxArea(height)
	}
}

// Benchmark with large input
func BenchmarkMaxAreaLarge(b *testing.B) {
	height := make([]int, 10000)
	for i := range height {
		height[i] = (i * 13) % 100
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxArea(height)
	}
}
