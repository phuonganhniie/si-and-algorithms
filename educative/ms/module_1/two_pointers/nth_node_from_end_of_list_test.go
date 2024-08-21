package twopointers

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/phuonganhniie/educative/helper"
)

func TestRemoveNthLastNode(t *testing.T) {
	tests := []struct {
		arr  []int
		n    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 1, []int{1, 2, 3, 4, 5, 6, 7, 9}},
		{[]int{32, 78, 65, 90, 12, 44}, 3, []int{32, 78, 65, 12, 44}},
		{[]int{23, 28, 10, 5, 67, 39, 70, 28}, 2, []int{23, 28, 10, 5, 67, 39, 28}},
	}

	countErr := 0
	for _, tt := range tests {
		ll := helper.BuildLinkedList(tt.arr)
		respLL := removeNthLastNode2(ll, tt.n)
		got := helper.BuildLinkedListToArray(respLL)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("removeNthLastNode test failed, want: %v - got: %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
