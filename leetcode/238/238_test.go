package leetcode_238

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{
			name: "basic",
			in:   []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
		},
		{
			name: "single_zero",
			in:   []int{1, 2, 0, 4},
			want: []int{0, 0, 8, 0},
		},
		{
			name: "multiple_zeros",
			in:   []int{0, 0, 3, 4},
			want: []int{0, 0, 0, 0},
		},
		{
			name: "negatives",
			in:   []int{-1, 2, -3, 4},
			want: []int{-24, 12, -8, 6},
		},
		{
			name: "two_elements",
			in:   []int{5, 6},
			want: []int{6, 5},
		},
		{
			name: "duplicates",
			in:   []int{2, 2, 2, 2},
			want: []int{8, 8, 8, 8},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ProductExceptSelf(tc.in)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("ProductExceptSelf(%v) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}
