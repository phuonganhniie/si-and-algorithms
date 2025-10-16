package leetcode_643

import (
	"fmt"
	"math"
	"testing"
)

func TestFindMaxAverage(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want float64
	}{
		{
			name: "Mixed positive and negative",
			nums: []int{-1, 5, -3, 7, -2, 4},
			k:    3,
			want: 3,
		},
		{
			name: "Single element array",
			nums: []int{5},
			k:    1,
			want: 5.00000,
		},
		{
			name: "Example 0: Array with 6 elements, k=4",
			nums: []int{1, 12, -5, -6, 50, 3},
			k:    4,
			want: 12.75,
		},
		{
			name: "Example 1: Array with 10 elements, k=7",
			nums: []int{7, 4, 5, 8, 8, 3, 9, 8, 7, 6},
			k:    7,
			want: 7.00000,
		},
		{
			name: "Example 2: Small array with k=2",
			nums: []int{4, 2, 1, 3, 3},
			k:    2,
			want: 3.00000,
		},
		{
			name: "Example 3: Array with negative numbers",
			nums: []int{1, 12, -5, -6, 50, 3},
			k:    4,
			want: 12.75000,
		},
		{
			name: "All same elements",
			nums: []int{3, 3, 3, 3, 3},
			k:    3,
			want: 3.00000,
		},
		{
			name: "k equals array length",
			nums: []int{1, 2, 3, 4, 5},
			k:    5,
			want: 3.00000,
		},
		{
			name: "Array with all negative numbers",
			nums: []int{-1, -2, -3, -4, -5},
			k:    2,
			want: -1.50000,
		},
		{
			name: "Maximum at the beginning",
			nums: []int{10, 9, 1, 2, 3, 4},
			k:    2,
			want: 9.50000,
		},
		{
			name: "Maximum at the end",
			nums: []int{1, 2, 3, 4, 9, 10},
			k:    2,
			want: 9.50000,
		},
		{
			name: "Maximum in the middle",
			nums: []int{1, 2, 8, 9, 3, 4},
			k:    2,
			want: 8.50000,
		},
		{
			name: "Large window size",
			nums: []int{0, 1, 1, 3, 3},
			k:    4,
			want: 2.00000,
		},
		{
			name: "Two elements, k=1",
			nums: []int{5, 3},
			k:    1,
			want: 5.00000,
		},
		{
			name: "Decimal result with larger window",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8},
			k:    4,
			want: 6.50000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMaxAverage(tt.nums, tt.k)
			// Use a small epsilon for floating point comparison
			if math.Abs(got-tt.want) > 0.00001 {
				t.Errorf("findMaxAverage() = %v, want %v", got, tt.want)
				fmt.Printf("Input: nums=%v, k=%d\n", tt.nums, tt.k)
			} else {
				fmt.Printf("✓ %s: PASSED (result: %.5f)\n", tt.name, got)
			}
		})
	}
}

// Benchmark test to measure performance
func BenchmarkFindMaxAverage(b *testing.B) {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	for i := 0; i < b.N; i++ {
		findMaxAverage(nums, k)
	}
}

// Benchmark with large input
func BenchmarkFindMaxAverageLarge(b *testing.B) {
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = i % 100
	}
	k := 100

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findMaxAverage(nums, k)
	}
}
