package leetcode_33

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "Example 1: Target found in right sorted portion",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
		{
			name:   "Example 2: Target not in array",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			want:   -1,
		},
		{
			name:   "Example 3: Single element not found",
			nums:   []int{1},
			target: 0,
			want:   -1,
		},
		{
			name:   "Example 4: Target in right portion",
			nums:   []int{4, 5, 6, 7, 0, 1, 2, 3},
			target: 1,
			want:   5,
		},
		{
			name:   "Single element found",
			nums:   []int{5},
			target: 5,
			want:   0,
		},
		{
			name:   "No rotation, target at beginning",
			nums:   []int{1, 2, 3, 4, 5, 6, 7},
			target: 1,
			want:   0,
		},
		{
			name:   "No rotation, target at end",
			nums:   []int{1, 2, 3, 4, 5, 6, 7},
			target: 7,
			want:   6,
		},
		{
			name:   "No rotation, target in middle",
			nums:   []int{1, 2, 3, 4, 5, 6, 7},
			target: 4,
			want:   3,
		},
		{
			name:   "No rotation, target not found",
			nums:   []int{1, 2, 3, 4, 5, 6, 7},
			target: 8,
			want:   -1,
		},
		{
			name:   "Rotated once, target in left portion",
			nums:   []int{7, 1, 2, 3, 4, 5, 6},
			target: 7,
			want:   0,
		},
		{
			name:   "Rotated once, target in right portion",
			nums:   []int{7, 1, 2, 3, 4, 5, 6},
			target: 3,
			want:   3,
		},
		{
			name:   "Rotated to last position, target found",
			nums:   []int{2, 3, 4, 5, 6, 7, 1},
			target: 1,
			want:   6,
		},
		{
			name:   "Rotated to last position, target at pivot",
			nums:   []int{2, 3, 4, 5, 6, 7, 1},
			target: 2,
			want:   0,
		},
		{
			name:   "Two elements, not rotated, target first",
			nums:   []int{1, 3},
			target: 1,
			want:   0,
		},
		{
			name:   "Two elements, not rotated, target second",
			nums:   []int{1, 3},
			target: 3,
			want:   1,
		},
		{
			name:   "Two elements, rotated, target first",
			nums:   []int{3, 1},
			target: 3,
			want:   0,
		},
		{
			name:   "Two elements, rotated, target second",
			nums:   []int{3, 1},
			target: 1,
			want:   1,
		},
		{
			name:   "Two elements, target not found",
			nums:   []int{3, 1},
			target: 2,
			want:   -1,
		},
		{
			name:   "Target at rotation point (minimum)",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
		{
			name:   "Target before rotation point",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 6,
			want:   2,
		},
		{
			name:   "Target after rotation point",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 2,
			want:   6,
		},
		{
			name:   "Large array, target in left portion",
			nums:   []int{10, 11, 12, 13, 14, 15, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			target: 12,
			want:   2,
		},
		{
			name:   "Large array, target in right portion",
			nums:   []int{10, 11, 12, 13, 14, 15, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			target: 5,
			want:   10,
		},
		{
			name:   "Large array, target at rotation point",
			nums:   []int{10, 11, 12, 13, 14, 15, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			target: 1,
			want:   6,
		},
		{
			name:   "Large array, target not found",
			nums:   []int{10, 11, 12, 13, 14, 15, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			target: 16,
			want:   -1,
		},
		{
			name:   "Negative numbers, rotated",
			nums:   []int{-2, -1, 0, 1, -10, -9, -8, -7, -6, -5, -4, -3},
			target: -9,
			want:   5,
		},
		{
			name:   "Negative numbers, target in left portion",
			nums:   []int{-2, -1, 0, 1, -10, -9, -8, -7, -6, -5, -4, -3},
			target: -1,
			want:   1,
		},
		{
			name:   "All negative, rotated",
			nums:   []int{-5, -4, -3, -2, -1, -10, -9, -8, -7, -6},
			target: -10,
			want:   5,
		},
		{
			name:   "Target is first element",
			nums:   []int{5, 6, 7, 0, 1, 2, 3, 4},
			target: 5,
			want:   0,
		},
		{
			name:   "Target is last element",
			nums:   []int{5, 6, 7, 0, 1, 2, 3, 4},
			target: 4,
			want:   7,
		},
		{
			name:   "Target is middle element",
			nums:   []int{5, 6, 7, 0, 1, 2, 3, 4},
			target: 0,
			want:   3,
		},
		{
			name:   "Rotated array with zero",
			nums:   []int{5, 6, 7, 8, 0, 1, 2, 3, 4},
			target: 0,
			want:   4,
		},
		{
			name:   "Large gap between elements",
			nums:   []int{1000, 2000, 3000, 100, 200, 300},
			target: 200,
			want:   4,
		},
		{
			name:   "Even length array, rotated",
			nums:   []int{4, 5, 6, 7, 0, 1, 2, 3},
			target: 7,
			want:   3,
		},
		{
			name:   "Odd length array, rotated",
			nums:   []int{5, 6, 7, 8, 9, 1, 2, 3, 4},
			target: 9,
			want:   4,
		},
		{
			name:   "Rotated in middle, target before pivot",
			nums:   []int{6, 7, 8, 9, 10, 1, 2, 3, 4, 5},
			target: 8,
			want:   2,
		},
		{
			name:   "Rotated in middle, target after pivot",
			nums:   []int{6, 7, 8, 9, 10, 1, 2, 3, 4, 5},
			target: 3,
			want:   7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := search(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
				fmt.Printf("Input: nums=%v, target=%d\n", tt.nums, tt.target)
			} else {
				fmt.Printf("✓ %s: PASSED (result: %d)\n", tt.name, got)
			}
		})
	}
}

// Benchmark test to measure performance
func BenchmarkSearch(b *testing.B) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	for i := 0; i < b.N; i++ {
		search(nums, target)
	}
}

// Benchmark with large input
func BenchmarkSearchLarge(b *testing.B) {
	nums := make([]int, 10000)
	rotatePoint := 7500

	// Create rotated sorted array
	for i := 0; i < len(nums); i++ {
		nums[i] = (i + rotatePoint) % len(nums)
	}
	target := 5000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		search(nums, target)
	}
}

// Benchmark target not found
func BenchmarkSearchNotFound(b *testing.B) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 10 // Not in array
	for i := 0; i < b.N; i++ {
		search(nums, target)
	}
}
