package arrays

import (
	"reflect"
	"testing"
)

func TestIndexArray(t *testing.T) {
	tests := []struct {
		arr  []int
		size int
		want []int
	}{
		{[]int{3, -1, 5, 6, -1, 4, 2}, 7, []int{-1, -1, 2, 3, 4, 5, 6}},
		{[]int{8, -1, 6, 1, 9, 3, 2, 7, 4, -1}, 10, []int{-1, 1, 2, 3, 4, -1, 6, 7, 8, 9}},
		{[]int{1, 3, 2, 4, -1, 0}, 6, []int{0, 1, 2, 3, 4, -1}},
		{[]int{3, -1, 2, 1, 4}, 5, []int{-1, 1, 2, 3, 4}},
	}

	for _, tt := range tests {
		got := IndexArray2(tt.arr, tt.size)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("IndexArray test failed, want %v - got %v", tt.want, got)
		}
	}
}
