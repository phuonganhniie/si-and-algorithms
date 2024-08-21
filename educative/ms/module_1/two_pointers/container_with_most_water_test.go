package twopointers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestContainerWithMostWater(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{[]int{7, 7, 2}, 7},
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{2, 8, 6, 3, 5, 4, 7}, 35},
		{[]int{1, 5}, 1},
		{[]int{1, 1}, 1},
	}

	countErr := 0
	for _, tt := range tests {
		got := containerWithMostWater(tt.height)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("containerWithMostWater test failed, want: %v - got: %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
