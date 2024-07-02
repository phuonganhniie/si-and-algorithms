package leetcode_350

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
	}

	countErr := 0
	for _, tt := range tests {
		got := intersect(tt.nums1, tt.nums2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("intersect test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
