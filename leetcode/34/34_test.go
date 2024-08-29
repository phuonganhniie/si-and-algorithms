package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{3, 9, 9, 9, 11, 12}, 9, []int{1, 3}},
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{}, 0, []int{-1, -1}},
	}

	countErr := 0
	for _, tt := range tests {
		got := searchRange(tt.nums, tt.target)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("searchRange test failed, want: %v - got: %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}

	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}

func BenchmarkSearchRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		searchRange([]int{3, 9, 9, 9, 11, 12}, 9)
	}
}
