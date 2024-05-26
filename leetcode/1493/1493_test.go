package leetcode_1493

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLongestSubarray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 1, 0, 1}, 3},
		{[]int{0, 1, 1, 1, 0, 1, 1, 0, 1}, 5},
		{[]int{1, 1, 1}, 2},
	}

	for _, tt := range tests {
		got := longestSubarray(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("longestSubarray test failed, want %v - got %v", tt.want, got)
		}
		fmt.Println("==================================")
	}
}
