package leetcode_1838

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxFrequency(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 4, 8, 13}, 5, 2},
		{[]int{1, 2, 4}, 5, 3},
		{[]int{3, 9, 6}, 2, 1},
	}

	countErr := 0
	for _, tt := range tests {
		got := maxFrequency(tt.nums, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("maxFrequency test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
