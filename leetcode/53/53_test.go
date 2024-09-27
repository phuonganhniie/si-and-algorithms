package leetcode_53

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{-4, -2, -3, -5}, -2},
		{[]int{4, -1, 2, 5, -2}, 10},
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
	}

	countErr := 0
	for i, tt := range tests {
		got := maxSubArray(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("maxSubArray test [%d] failed, want %v - got %v", i+1, tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
