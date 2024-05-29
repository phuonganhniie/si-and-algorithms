package leetcode_1877

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMinPairSum(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 5, 2, 3}, 7},
		{[]int{3, 5, 4, 2, 4, 6}, 8},
	}

	countErr := 0
	for _, tt := range tests {
		got := minPairSum(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("minPairSum test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
