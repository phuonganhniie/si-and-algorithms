package sorting

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{[]int{8, 2, 4, 7, 1, 3, 9, 6, 5}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	for _, tt := range tests {
		got := QuickSort(tt.arr)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("QuickSort test failed, want %v - got %v", tt.want, got)
		}
	}
}
