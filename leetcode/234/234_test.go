package leetcode_234

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/phuonganhniie/leetcode/helper"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		head []int
		want bool
	}{
		{[]int{1, 2, 2, 1}, true},
		{[]int{1, 2}, false},
	}

	countErr := 0
	for _, tt := range tests {
		head := helper.BuildLinkedList(tt.head)
		got := isPalindrome(head)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("isPalindrome test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
