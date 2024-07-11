package twopointers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		colors []int
		want   []int
	}{
		{[]int{0, 1, 0}, []int{0, 0, 1}},
		{[]int{1}, []int{1}},
		{[]int{2, 2}, []int{2, 2}},
		{[]int{1, 1, 0, 2}, []int{0, 1, 1, 2}},
		{[]int{2, 1, 1, 0, 0}, []int{0, 0, 1, 1, 2}},
	}

	countErr := 0
	for _, tt := range tests {
		got := sortColors(tt.colors)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("sortColors test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
