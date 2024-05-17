package arrays

import (
	"reflect"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		nums  []int
		value int
		want  bool
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 5, true},
		{[]int{5, 6, 7, 8, 9, 10}, 8, true},
		{[]int{1, 8, 9, 10, 50, 108}, 10, true},
		{[]int{101, 201, 301, 401}, 501, false},
	}

	for _, tt := range tests {
		got := BinarySearch(tt.nums, tt.value)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("BinarySearch test failed, want %v - got %v", tt.want, got)
		}
	}
}
