package leetcode_2215

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindDifference(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  [][]int
	}{
		{[]int{-80, -15, -81, -28, -61, 63, 14, -45, -35, -10}, []int{-1, -40, -44, 41, 10, -43, 69, 10, 2}, [][]int{{-81, -35, -10, -28, -61, -45, -15, 14, -80, 63}, {-1, 2, 69, -40, 41, 10, -43, -44}}},
		{[]int{1, 2, 3}, []int{2, 4, 6}, [][]int{{1, 3}, {4, 6}}},
		{[]int{1, 2, 3, 3}, []int{1, 1, 2, 2}, [][]int{{3}, {}}},
	}

	for _, tt := range tests {
		got := findDifference(tt.nums1, tt.nums2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("findDifference test failed, want %v - got %v", tt.want, got)
		}
		fmt.Println("==================================")
	}
}
