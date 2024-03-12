package leetcode_1171

import "testing"

func sliceToList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return dummy.Next
}

func compareLists(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}

func TestRemoveZeroSumSublists(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, -3, 3, 1}, []int{3, 1}},
		{[]int{1, 2, 3, -3, 4}, []int{1, 2, 4}},
		{[]int{1, 2, 3, -3, -2}, []int{1}},
	}

	for _, tt := range tests {
		head := sliceToList(tt.input)
		got := RemoveZeroSumSublists(head)
		want := sliceToList(tt.want)
		if !compareLists(got, want) {
			t.Errorf("removeZeroSumSublists(%v) = %v, want %v", tt.input, got, want)
		}
	}
}
