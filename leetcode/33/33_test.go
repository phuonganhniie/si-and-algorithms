package leetcode_33

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
		{[]int{4, 5, 6, 7, 0, 1, 2, 3}, 1, 5},
	}

	countErr := 0

	for _, tt := range cases {
		got := search(tt.nums, tt.target)
		if tt.want != got {
			t.Errorf("search test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}

	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
