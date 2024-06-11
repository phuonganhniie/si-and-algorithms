package leetcode_523

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCheckSubarraySum(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want bool
	}{
		{[]int{5, 0, 0, 0}, 3, true},
		{[]int{23, 2, 4, 6, 6}, 7, true},
		{[]int{23, 2, 6, 4, 7}, 6, true},
		{[]int{23, 2, 4, 6, 7}, 6, true},
		{[]int{23, 2, 6, 4, 7}, 13, false},
	}

	countErr := 0
	for _, tt := range tests {
		got := checkSubarraySum(tt.nums, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("checkSubarraySum test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
