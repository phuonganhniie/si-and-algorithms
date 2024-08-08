package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSmallestPositiveMissingNumber(t *testing.T) {
	tests := []struct {
		arr  []int
		size int
		want int
	}{
		{[]int{8, 5, 9, 3, 2, 7, 6, 4, 10}, 9, 1},
		{[]int{1, 3, 4, 2, 6}, 5, 5},
		{[]int{5, 2, 1, 4}, 4, 3},
		{[]int{3, 7, 5, 6, 1, 4, 2}, 7, -1},
	}

	countErr := 0
	for _, tt := range tests {
		got := smallestPositiveMissingNumber(tt.arr, tt.size)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("smallestPositiveMissingNumber test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
