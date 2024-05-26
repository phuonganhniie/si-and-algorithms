package leetcode_1004

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLongestOnes(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
		{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3, 10},
	}

	for _, tt := range tests {
		got := longestOnes(tt.nums, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("longestOnes test failed, want %v - got %v", tt.want, got)
		}
		fmt.Println("==================================")
	}
}
