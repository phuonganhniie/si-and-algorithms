package leetcode_21

import (
	"github.com/phuonganhniie/leetcode/helper"
)

func MergeTwoLists(list1 *helper.ListNode, list2 *helper.ListNode) *helper.ListNode {
	mergeList := &helper.ListNode{}
	current := mergeList

	for ; list1 != nil && list2 != nil; current = current.Next {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
	}

	if list1 == nil {
		current.Next = list2
	} else {
		current.Next = list1
	}

	return mergeList.Next
}
