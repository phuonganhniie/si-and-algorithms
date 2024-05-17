package arrays

import (
	"reflect"
	"testing"
)

func TestRotateArray(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 2, []int{3, 4, 5, 6, 1, 2}},
		{[]int{7, 8, 9, 10}, 3, []int{10, 7, 8, 9}},
		{[]int{3, 4, 5, 6, 7, 8, 9}, 5, []int{8, 9, 3, 4, 5, 6, 7}},
		{[]int{7, 3, 9, 10, 3}, 4, []int{3, 7, 3, 9, 10}},
	}

	for _, tt := range tests {
		got := RotateArray(tt.nums, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("RotateArray test failed, want %v - got %v", tt.want, got)
		}
	}
}
