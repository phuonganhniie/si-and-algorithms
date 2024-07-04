package leetcode_2181

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/phuonganhniie/leetcode/helper"
)

func TestMergeNodes(t *testing.T) {
	tests := []struct {
		head []int
		want []int
	}{
		{[]int{0, 3, 1, 0, 4, 5, 2, 0}, []int{4, 11}},
		{[]int{0, 1, 0, 3, 0, 2, 2, 0}, []int{1, 3, 4}},
	}

	countErr := 0
	for _, tt := range tests {
		listNodes := helper.BuildLinkedList(tt.head)
		res := mergeNodes(listNodes)
		got := helper.BuildLinkedListToArray(res)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("mergeNodes test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
