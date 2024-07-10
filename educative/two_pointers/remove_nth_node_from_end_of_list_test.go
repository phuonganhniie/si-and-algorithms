package twopointers

import (
	"reflect"
	"testing"

	"github.com/phuonganhniie/educative/helper"
)

func TestRemoveNthLastNode(t *testing.T) {
	tests := []struct {
		head []int
		n    int
		want []int
	}{
		{[]int{23, 28, 10, 5, 67, 39, 70, 28}, 2, []int{23, 28, 10, 5, 67, 39, 28}},
		{[]int{34, 53, 6, 95, 38, 28, 17, 63, 16, 76}, 3, []int{34, 53, 6, 95, 38, 28, 17, 16, 76}},
		{[]int{288, 224, 275, 390, 4, 383, 330, 60, 193}, 4, []int{288, 224, 275, 390, 4, 330, 60, 193}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 1, []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{[]int{69, 8, 49, 106, 116, 112}, 6, []int{8, 49, 106, 116, 112}},
	}

	for _, tt := range tests {
		head := helper.BuildLinkedList(tt.head)
		got := removeNthLastNode(head, tt.n)
		gotList := helper.BuildLinkedListToArray(got)

		if !reflect.DeepEqual(gotList, tt.want) {
			t.Errorf("removeNthLastNode test failed, want %v - got %v", tt.want, got)
		}
	}
}
