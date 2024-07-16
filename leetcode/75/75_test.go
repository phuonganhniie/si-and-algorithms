package leetcode_75

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 0, 2, 1, 1, 0, 2, 2, 0}, []int{0, 0, 0, 1, 1, 1, 2, 2, 2}},
		{[]int{1, 2, 0}, []int{0, 1, 2}},
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{[]int{2, 0, 1}, []int{0, 1, 2}},
	}

	countErr := 0
	for _, tt := range tests {
		got := sortColors2(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("sortColors test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
