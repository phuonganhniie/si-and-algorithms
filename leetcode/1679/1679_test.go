package leetcode_1679

import (
	"reflect"
	"testing"
)

func TestMaxOperations(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{4, 4, 1, 3, 1, 3, 2, 2, 5, 5, 1, 5, 2, 1, 2, 3, 5, 4}, 2, 2},
		{[]int{1, 2, 3, 4}, 5, 2},
		{[]int{3, 1, 3, 4, 3}, 6, 1},
	}

	for _, tt := range tests {
		got := maxOperations(tt.nums, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("maxOperations test failed, want %v - got %v", tt.want, got)
		}
	}
}
