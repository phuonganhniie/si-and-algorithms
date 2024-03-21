package leetcode_238

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}

	for _, tt := range tests {
		got := ProductExceptSelf(tt.input)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("ProductExceptSelf(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}
