package binarysearch

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
		{[]int{6, 7, 1, 2, 3, 4, 5}, 6, 0},
		{[]int{6, 7, 1, 2, 3, 4, 5}, 3, 4},
	}

	for _, tt := range tests {
		got := binarySearchRotated(tt.nums, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("binarySearchRotated test failed, want %v - got %v", tt.want, got)
		}
		fmt.Println("==================================")
	}
}
