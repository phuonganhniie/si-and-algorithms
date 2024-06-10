package leetcode_1051

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHeightChecker(t *testing.T) {
	tests := []struct {
		heights []int
		want    int
	}{
		{[]int{1, 1, 4, 2, 1, 3}, 3},
		{[]int{5, 1, 2, 3, 4}, 5},
		{[]int{1, 2, 3, 4, 5}, 0},
	}

	countErr := 0
	for _, tt := range tests {
		got := heightCheckerCountingSort(tt.heights)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("heightChecker test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
