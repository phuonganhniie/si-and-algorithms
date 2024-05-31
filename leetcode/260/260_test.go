package leetcode_260

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{[]int{1, 2, 1, 3, 2, 5}, []int{3, 5}},
		{[]int{-1, 0}, []int{-1, 0}},
		{[]int{0, 1}, []int{1, 0}},
	}

	countErr := 0
	for _, tt := range tests {
		got := singleNumberOptimize(tt.arr)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("singleNumber test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
