package leetcode_121

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "Example 1: Buy on day 1, sell on day 5",
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			name:   "Example 2: No profit possible (decreasing prices)",
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
		{
			name:   "Single element array",
			prices: []int{5},
			want:   0,
		},
		{
			name:   "Two elements with profit",
			prices: []int{1, 5},
			want:   4,
		},
		{
			name:   "Two elements with no profit",
			prices: []int{5, 1},
			want:   0,
		},
		{
			name:   "All same prices",
			prices: []int{3, 3, 3, 3, 3},
			want:   0,
		},
		{
			name:   "Maximum profit at the end",
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
		},
		{
			name:   "Lowest price at the end",
			prices: []int{5, 4, 3, 2, 1},
			want:   0,
		},
		{
			name:   "Buy at minimum, sell at maximum",
			prices: []int{3, 2, 6, 5, 0, 3},
			want:   4,
		},
		{
			name:   "Multiple peaks and valleys",
			prices: []int{2, 4, 1, 7, 5, 11},
			want:   10,
		},
		{
			name:   "Profit at the beginning",
			prices: []int{1, 10, 5, 3, 2},
			want:   9,
		},
		{
			name:   "Large price difference",
			prices: []int{100, 1, 101, 50},
			want:   100,
		},
		{
			name:   "Small fluctuations",
			prices: []int{1, 2, 1, 2, 1, 2},
			want:   1,
		},
		{
			name:   "V-shaped pattern",
			prices: []int{5, 3, 1, 3, 5},
			want:   4,
		},
		{
			name:   "Inverted V-shaped pattern",
			prices: []int{1, 3, 5, 3, 1},
			want:   4,
		},
		{
			name:   "Zero in prices",
			prices: []int{0, 5, 2, 8},
			want:   8,
		},
		{
			name:   "Large array with best profit in middle",
			prices: []int{10, 9, 8, 1, 15, 2, 3},
			want:   14,
		},
		{
			name:   "Minimum at start, maximum at end",
			prices: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:   9,
		},
		{
			name:   "Early high, late recovery",
			prices: []int{10, 5, 1, 2, 3, 9},
			want:   8,
		},
		{
			name:   "Multiple small gains",
			prices: []int{1, 2, 3, 2, 3, 4},
			want:   3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit(tt.prices)
			if got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
				fmt.Printf("Input: %v\n", tt.prices)
			} else {
				fmt.Printf("✓ %s: PASSED (result: %d)\n", tt.name, got)
			}
		})
	}
}

// Benchmark test to measure performance
func BenchmarkMaxProfit(b *testing.B) {
	prices := []int{7, 1, 5, 3, 6, 4}
	for i := 0; i < b.N; i++ {
		maxProfit(prices)
	}
}

// Benchmark with large input
func BenchmarkMaxProfitLarge(b *testing.B) {
	prices := make([]int, 10000)
	for i := range prices {
		prices[i] = (i*7 + 13) % 100
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxProfit(prices)
	}
}
