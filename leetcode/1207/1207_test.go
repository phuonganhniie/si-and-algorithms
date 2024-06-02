package leetcode_1207

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUniqueOccurrences(t *testing.T) {
	tests := []struct {
		arr  []int
		want bool
	}{
		{[]int{-1, -1, 1, 1, 2, 2, 3}, false},
		{[]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}, true},
		{[]int{1, 2, 2, 1, 1, 3}, true},
		{[]int{1, 2}, false},
		{[]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, true},
		{[]int{}, true},
		{[]int{1}, true},
		{[]int{1, 1, 1, 1}, true},
		{[]int{1, 2, 2, 3, 3}, false},
		{[]int{0, 0, 1, 1, 1}, true},
	}

	countErr := 0
	for i, tt := range tests {
		got := uniqueOccurrences(tt.arr)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("uniqueOccurrences test case %v failed, want %v - got %v", i+1, tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
