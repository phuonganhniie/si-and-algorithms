package leetcode_1669

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

func TestMergeInBetween(t *testing.T) {
	tests := []struct {
		list1 []int
		a, b  int
		list2 []int
		want  []int
	}{
		{[]int{10, 1, 13, 6, 9, 5}, 3, 4, []int{1000000, 1000001, 1000002}, []int{10, 1, 13, 1000000, 1000001, 1000002, 5}},
		{[]int{0, 1, 2, 3, 4, 5, 6}, 2, 5, []int{1000000, 1000001, 1000002, 1000003, 1000004}, []int{0, 1, 1000000, 1000001, 1000002, 1000003, 1000004, 6}},
	}

	for _, tt := range tests {
		list1 := createList(tt.list1)
		list2 := createList(tt.list2)
		got := MergeInBetween(list1, tt.a, tt.b, list2)
		if !reflect.DeepEqual(listToSlice(got), tt.want) {
			t.Errorf("mergeInBetween(%v, %d, %d, %v) = %v, want %v", tt.list1, tt.a, tt.b, tt.list2, listToSlice(got), tt.want)
		}
	}
}
