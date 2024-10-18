package leetcode_162

import (
	"fmt"
	"testing"
)

func TestFindPeakElement(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{5, 4, 3, 4, 5}, 0},
		{[]int{1, 2, 1, 3, 5, 6, 4}, 5},
		{[]int{2, 1}, 0},
		{[]int{1}, 0},
		{[]int{1, 2, 3, 1}, 2},
	}

	countErr := 0

	for _, c := range cases {
		got := findPeakElement(c.nums)
		if got != c.want {
			t.Errorf("findPeakElement test failed, want %v - got %v", c.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}

	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
