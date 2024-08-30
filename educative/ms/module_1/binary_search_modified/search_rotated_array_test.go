package binarysearchmodified

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBinarySearchRotated(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{5, 1, 3}, 3, 2},
		{[]int{6, 7, 1, 2, 3, 4, 5}, 3, 4},
		{[]int{6, 7, 1, 2, 3, 4, 5}, 6, 0},
		{[]int{4, 5, 6, 1, 2, 3}, 3, 5},
		{[]int{4}, 1, -1},
	}

	countErr := 0

	for _, tt := range tests {
		got := binarySearchRotated(tt.nums, tt.target)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("binarySearchRotated test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}

	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}

func BenchmarkBinarySearchRotated(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binarySearchRotated([]int{6, 7, 1, 2, 3, 4, 5}, 3)
	}
}
