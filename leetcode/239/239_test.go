package leetcode_239

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{2, 4, 3, 6, 5, 4, 1, 10}, 3, []int{4, 6, 6, 6, 5, 10}},
		{[]int{1, -1}, 1, []int{1, -1}},
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{[]int{1}, 1, []int{1}},
	}

	countErr := 0
	for _, tt := range tests {
		got := maxSlidingWindowOptimize(tt.nums, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("maxSlidingWindowOptimize test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
