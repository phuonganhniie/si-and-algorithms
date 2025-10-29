package leetcode_704

import (
	"fmt"
	"testing"
)

func TestSearchBasic(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "Example 1: Target found in middle-right",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			want:   4,
		},
		{
			name:   "Example 2: Target not in array",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			want:   -1,
		},
		{
			name:   "Single element found",
			nums:   []int{5},
			target: 5,
			want:   0,
		},
		{
			name:   "Single element not found",
			nums:   []int{5},
			target: 3,
			want:   -1,
		},
		{
			name:   "Target at first position",
			nums:   []int{1, 2, 3, 4, 5},
			target: 1,
			want:   0,
		},
		{
			name:   "Target at last position",
			nums:   []int{1, 2, 3, 4, 5},
			target: 5,
			want:   4,
		},
		{
			name:   "Target in the middle",
			nums:   []int{1, 2, 3, 4, 5},
			target: 3,
			want:   2,
		},
		{
			name:   "Two elements, target is first",
			nums:   []int{1, 3},
			target: 1,
			want:   0,
		},
		{
			name:   "Two elements, target is second",
			nums:   []int{1, 3},
			target: 3,
			want:   1,
		},
		{
			name:   "Two elements, target not found (too small)",
			nums:   []int{1, 3},
			target: 0,
			want:   -1,
		},
		{
			name:   "Two elements, target not found (too large)",
			nums:   []int{1, 3},
			target: 5,
			want:   -1,
		},
		{
			name:   "Two elements, target not found (in between)",
			nums:   []int{1, 3},
			target: 2,
			want:   -1,
		},
		{
			name:   "Target smaller than all elements",
			nums:   []int{10, 20, 30, 40, 50},
			target: 5,
			want:   -1,
		},
		{
			name:   "Target larger than all elements",
			nums:   []int{10, 20, 30, 40, 50},
			target: 60,
			want:   -1,
		},
		{
			name:   "Large array with negative numbers",
			nums:   []int{-100, -50, -10, 0, 10, 50, 100},
			target: -10,
			want:   2,
		},
		{
			name:   "All negative numbers",
			nums:   []int{-5, -4, -3, -2, -1},
			target: -3,
			want:   2,
		},
		{
			name:   "All same numbers (edge case)",
			nums:   []int{5, 5, 5, 5, 5},
			target: 5,
			want:   2, // Could be any valid index, binary search returns one
		},
		{
			name:   "Large array, target at beginning",
			nums:   []int{1, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
			target: 1,
			want:   0,
		},
		{
			name:   "Large array, target at end",
			nums:   []int{1, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
			target: 100,
			want:   10,
		},
		{
			name:   "Large array, target in middle",
			nums:   []int{1, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
			target: 50,
			want:   5,
		},
		{
			name:   "Large array, target not found",
			nums:   []int{1, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
			target: 45,
			want:   -1,
		},
		{
			name:   "Even length array",
			nums:   []int{2, 4, 6, 8},
			target: 6,
			want:   2,
		},
		{
			name:   "Odd length array",
			nums:   []int{2, 4, 6, 8, 10},
			target: 6,
			want:   2,
		},
		{
			name:   "Sequential numbers from 0",
			nums:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			target: 7,
			want:   7,
		},
		{
			name:   "Powers of 2",
			nums:   []int{1, 2, 4, 8, 16, 32, 64, 128, 256},
			target: 64,
			want:   6,
		},
		{
			name:   "Large sparse array",
			nums:   []int{1, 100, 200, 300, 400, 500, 600, 700, 800, 900, 1000},
			target: 700,
			want:   7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BasicSearch(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("BasicSearch() = %v, want %v", got, tt.want)
				fmt.Printf("Input: nums=%v, target=%d\n", tt.nums, tt.target)
			} else {
				fmt.Printf("✓ %s: PASSED (result: %d)\n", tt.name, got)
			}
		})
	}
}

// Benchmark test to measure performance
func BenchmarkBasicSearch(b *testing.B) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	for i := 0; i < b.N; i++ {
		BasicSearch(nums, target)
	}
}

// Benchmark with large input
func BenchmarkBasicSearchLarge(b *testing.B) {
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = i * 2
	}
	target := 9998

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BasicSearch(nums, target)
	}
}

// Benchmark worst case - target not found
func BenchmarkBasicSearchNotFound(b *testing.B) {
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = i * 2
	}
	target := 9999 // Odd number, won't be found

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BasicSearch(nums, target)
	}
}
