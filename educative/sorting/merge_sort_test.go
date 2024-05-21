package sorting

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{[]int{3, 7, 8, 5, 4, 2, 6, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8}},
	}

	for _, tt := range tests {
		got := MergeSort(tt.arr)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("MergeSort test failed, want %v - got %v", tt.want, got)
		}
	}
}
