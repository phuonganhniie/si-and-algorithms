package leetcode_283

import (
	"reflect"
	"testing"
)

func TestMoveZeros(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0}, []int{0}},
	}

	for _, tt := range tests {
		got := moveZeroes(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("moveZeroes test failed, want %v - got %v", tt.want, got)
		}
	}
}
