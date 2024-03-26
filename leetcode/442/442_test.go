package leetcode_442

import (
	"reflect"
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{2, 3}},
		{[]int{1, 1, 2}, []int{1}},
		{[]int{1}, []int{}},
	}

	for _, tt := range tests {
		got := FindDuplicates2(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("FindDuplicates(%v) = %d, want %d", tt.nums, got, tt.want)
		}
	}
}
