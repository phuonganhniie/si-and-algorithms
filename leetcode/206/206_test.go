package leetcode_206

import (
	"reflect"
	"testing"
)

func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	current := head

	for _, num := range nums[1:] {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}

	return head
}

func listToSlice(head *ListNode) []int {
	var nums []int

	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	return nums
}

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
		list := createList(tt.list)
		got := ReverseList(list)
		if !reflect.DeepEqual(listToSlice(got), tt.want) {
			t.Errorf("reverseList = %v, want %v", listToSlice(got), tt.want)
		}
	}
}
