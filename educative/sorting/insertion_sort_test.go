package sorting

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5, 0}, []int{0, 1, 2, 3, 4, 5}},
		{[]int{9, 1, 8, 2, 7, 3, 6, 4, 5}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	for _, tt := range tests {
		got := InsertionSort(tt.arr)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("InsertionSort test failed, want %v - got %v", tt.want, got)
		}
	}
}
