package leetcode_206

import (
	"reflect"
	"testing"

	"github.com/phuonganhniie/leetcode/helper"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		list []int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{}, []int{}},
	}

	for _, tt := range tests {
		head := helper.BuildLinkedList(tt.list)
		got := ReverseList(head)
		if !reflect.DeepEqual(helper.BuildLinkedListToArray(got), tt.want) {
			t.Errorf("reverseList got %v, want %v", helper.BuildLinkedListToArray(got), tt.want)
		}
	}
}
