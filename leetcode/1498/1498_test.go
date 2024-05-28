package leetcode_1498

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNumSubseq(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{27, 21, 14, 2, 15, 1, 19, 8, 12, 24, 21, 8, 12, 10, 11, 30, 15, 18, 28, 14, 26, 9, 2, 24, 23, 11, 7, 12, 9, 17, 30, 9, 28, 2, 14, 22, 19, 19, 27, 6, 15, 12, 29, 2, 30, 11, 20, 30, 21, 20, 2, 22, 6, 14, 13, 19, 21, 10, 18, 30, 2, 20, 28, 22}, 31, 688052206},
		{[]int{3, 5, 6, 7}, 9, 4},
		{[]int{3, 3, 6, 8}, 10, 6},
		{[]int{2, 3, 3, 4, 6, 7}, 12, 61},
	}

	countErr := 0
	for _, tt := range tests {
		got := numSubseq(tt.nums, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("numSubseq test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
