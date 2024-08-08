package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestArrayIndexMaxDiff(t *testing.T) {
	tests := []struct {
		arr  []int
		size int
		want int
	}{
		{[]int{3, 5, 2, 1, 4}, 5, 4},
		{[]int{8, 5, 6, 9, 3, 2, 7, 4, 10}, 9, 8},
		{[]int{1, 3, 4, 2, 6}, 5, 4},
		{[]int{3, 7, 5, 6, 1, 4, 2}, 7, 5},
	}

	countErr := 0
	for _, tt := range tests {
		got := arrayIndexMaxDiff(tt.arr, tt.size)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("arrayIndexMaxDiff test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
