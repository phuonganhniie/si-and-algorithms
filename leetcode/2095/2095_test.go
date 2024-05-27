package leetcode_2095

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/phuonganhniie/leetcode/helper"
)

func TestDeleteMiddle(t *testing.T) {
	tests := []struct {
		head []int
		want []int
	}{
		{[]int{1, 3, 4, 7, 1, 2, 6}, []int{1, 3, 4, 1, 2, 6}},
		{[]int{1, 2, 3, 4}, []int{1, 2, 4}},
		{[]int{2, 1}, []int{2}},
	}

	for _, tt := range tests {
		head := helper.BuildLinkedList(tt.head)
		got := deleteMiddle(head)
		if !reflect.DeepEqual(helper.BuildLinkedListToArray(got), tt.want) {
			t.Errorf("binarySearchRotated test failed, want %v - got %v", tt.want, helper.BuildLinkedListToArray(got))
		}
		fmt.Println("==================================")
	}
}
