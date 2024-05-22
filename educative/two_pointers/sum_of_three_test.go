package twopointers

import (
	"reflect"
	"testing"
)

func TestFindSumOfThree(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   bool
	}{
		{[]int{-1, 2, 1, -4, 5, -3}, -8, true},
		{[]int{3, 7, 1, 2, 8, 4, 5}, 10, true},
		{[]int{1, -1, 0}, -1, false},
		{[]int{3, 7, 1, 2, 8, 4, 5}, 21, false},
	}

	for _, tt := range tests {
		got := findSumOfThree(tt.nums, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("findSumOfThree test failed, want %v - got %v", tt.want, got)
		}
	}
}
