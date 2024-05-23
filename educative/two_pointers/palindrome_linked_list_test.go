package twopointers

import (
	"reflect"
	"testing"

	"github.com/phuonganhniie/leetcode/helper"
)

func TestPalindrome(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{4, 7, 9, 5, 4}, false},
		{[]int{3, 4, 5, 6, 5, 4, 3}, true},
		{[]int{1, 2, 3, 2, 1}, true},
		{[]int{6, 1, 0, 5, 1, 6}, false},
		{[]int{1, 2}, false},
	}

	for i, tt := range tests {
		head := helper.BuildLinkedList(tt.nums)
		got := palindrome(head)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("palindrome test case %v failed, want %v - got %v", i+1, tt.want, got)
		}
	}
}
