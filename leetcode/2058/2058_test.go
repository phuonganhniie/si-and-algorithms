package leetcode_2058

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/phuonganhniie/leetcode/helper"
)

func TestNodesBetweenCriticalPoints(t *testing.T) {
	tests := []struct {
		head []int
		want []int
	}{
		{[]int{3, 1}, []int{-1, -1}},
		{[]int{5, 3, 1, 2, 5, 1, 2}, []int{1, 3}},
		{[]int{1, 3, 2, 2, 3, 2, 2, 2, 7}, []int{3, 3}},
	}

	countErr := 0
	for _, tt := range tests {
		list := helper.BuildLinkedList(tt.head)
		got := nodesBetweenCriticalPoints(list)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("nodesBetweenCriticalPoints test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
