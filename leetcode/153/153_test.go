package leetcode_153

import (
	"fmt"
	"testing"
)

func TestFindMin(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1: Array rotated 3 times",
			nums: []int{3, 4, 5, 1, 2},
			want: 1,
		},
		{
			name: "Example 2: Array rotated 4 times",
			nums: []int{4, 5, 6, 7, 0, 1, 2},
			want: 0,
		},
		{
			name: "Example 3: Array rotated 10 times",
			nums: []int{11, 13, 15, 17},
			want: 11,
		},
		{
			name: "Single element",
			nums: []int{5},
			want: 5,
		},
		{
			name: "Two elements, not rotated",
			nums: []int{1, 2},
			want: 1,
		},
		{
			name: "Two elements, rotated once",
			nums: []int{2, 1},
			want: 1,
		},
		{
			name: "Three elements, not rotated",
			nums: []int{1, 2, 3},
			want: 1,
		},
		{
			name: "Three elements, rotated once",
			nums: []int{2, 3, 1},
			want: 1,
		},
		{
			name: "Three elements, rotated twice",
			nums: []int{3, 1, 2},
			want: 1,
		},
		{
			name: "No rotation (sorted array)",
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			want: 1,
		},
		{
			name: "Rotated once",
			nums: []int{7, 1, 2, 3, 4, 5, 6},
			want: 1,
		},
		{
			name: "Rotated to last position",
			nums: []int{2, 3, 4, 5, 6, 7, 1},
			want: 1,
		},
		{
			name: "Minimum at beginning",
			nums: []int{1, 2, 3, 4, 5},
			want: 1,
		},
		{
			name: "Minimum at end",
			nums: []int{2, 3, 4, 5, 1},
			want: 1,
		},
		{
			name: "Minimum in middle",
			nums: []int{4, 5, 1, 2, 3},
			want: 1,
		},
		{
			name: "Large values",
			nums: []int{100, 200, 300, 10, 20, 30},
			want: 10,
		},
		{
			name: "Negative numbers",
			nums: []int{-5, -3, -1, -10, -8, -6},
			want: -10,
		},
		{
			name: "Mixed positive and negative",
			nums: []int{3, 4, 5, -1, 0, 1, 2},
			want: -1,
		},
		{
			name: "All negative, rotated",
			nums: []int{-2, -1, -10, -9, -8, -7, -6, -5, -4, -3},
			want: -10,
		},
		{
			name: "Large array, minimum near start",
			nums: []int{10, 11, 12, 13, 14, 15, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: 1,
		},
		{
			name: "Large array, minimum near end",
			nums: []int{8, 9, 10, 11, 12, 13, 14, 15, 1, 2, 3, 4, 5, 6, 7},
			want: 1,
		},
		{
			name: "Large array, minimum in middle",
			nums: []int{5, 6, 7, 8, 9, 10, 11, 12, 1, 2, 3, 4},
			want: 1,
		},
		{
			name: "Consecutive numbers starting from 0",
			nums: []int{3, 4, 5, 6, 7, 0, 1, 2},
			want: 0,
		},
		{
			name: "Large gap between elements",
			nums: []int{1000, 2000, 3000, 100, 200, 300},
			want: 100,
		},
		{
			name: "Small rotation (1 step)",
			nums: []int{2, 1},
			want: 1,
		},
		{
			name: "Array with duplicates at edges (all unique)",
			nums: []int{10, 1, 10},
			want: 1,
		},
		{
			name: "Rotated array with zero",
			nums: []int{5, 6, 7, 0, 1, 2, 3, 4},
			want: 0,
		},
		{
			name: "Large sorted array (no rotation)",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			want: 1,
		},
		{
			name: "Even length array, rotated",
			nums: []int{4, 5, 6, 7, 1, 2, 3},
			want: 1,
		},
		{
			name: "Odd length array, rotated",
			nums: []int{5, 6, 7, 8, 1, 2, 3, 4},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMin(tt.nums)
			if got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
				fmt.Printf("Input: %v\n", tt.nums)
			} else {
				fmt.Printf("✓ %s: PASSED (result: %d)\n", tt.name, got)
			}
		})
	}
}

// Benchmark test to measure performance
func BenchmarkFindMin(b *testing.B) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	for i := 0; i < b.N; i++ {
		findMin(nums)
	}
}

// Benchmark with large input
func BenchmarkFindMinLarge(b *testing.B) {
	nums := make([]int, 10000)
	rotatePoint := 7500

	// Create rotated sorted array
	for i := 0; i < len(nums); i++ {
		nums[i] = (i + rotatePoint) % len(nums)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findMin(nums)
	}
}

// Benchmark no rotation case
func BenchmarkFindMinNoRotation(b *testing.B) {
	nums := make([]int, 10000)
	for i := 0; i < len(nums); i++ {
		nums[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findMin(nums)
	}
}
