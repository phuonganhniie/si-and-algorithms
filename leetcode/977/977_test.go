package leetcode_977

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
	}

	countErr := 0
	for i, tt := range tests {
		got := sortedSquares(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("sortedSquares test case %v failed, want %v - got %v", i+1, tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
