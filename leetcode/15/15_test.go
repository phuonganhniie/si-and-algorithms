package leetcode_15

import (
	"fmt"
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 1, 1}, [][]int{}},
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
	}

	for _, tt := range tests {
		got := threeSum(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("threeSum test failed, want %v - got %v", tt.want, got)
		}
		fmt.Println("==================================")
	}
}
