package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindGCD(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 5, 6, 9, 10}, 2},
		{[]int{7, 5, 6, 8, 3}, 1},
	}

	countErr := 0
	for _, tt := range tests {
		got := findGCD(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("findGCD test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
