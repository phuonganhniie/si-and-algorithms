package leetcode_643

import (
	"reflect"
	"testing"
)

func TestFindMaxAverage(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want float64
	}{
		{[]int{7, 4, 5, 8, 8, 3, 9, 8, 7, 6}, 7, 7.00000},
		{[]int{4, 2, 1, 3, 3}, 2, 3.00000},
		{[]int{1, 12, -5, -6, 50, 3}, 4, 12.75000},
		{[]int{5}, 1, 5.00000},
	}

	for _, tt := range tests {
		got := findMaxAverage(tt.nums, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("findMaxAverage test failed, want %v - got %v", tt.want, got)
		}
	}
}
