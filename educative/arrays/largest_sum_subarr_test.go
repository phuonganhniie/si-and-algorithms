package arrays

import (
	"reflect"
	"testing"
)

func TestMaxSubArraySum(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, -2, 3, 4, -4, 6, -14, 6, 2}, 9},
		{[]int{1, 7, -2, -5, 10, -1}, 11},
		{[]int{-5, 6, 3, -8, 2, 9}, 12},
		{[]int{-11, 8, 1, -3, 7}, 13},
	}

	for _, tt := range tests {
		got := MaxSubArraySum(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("MaxSubArraySum test failed, want %v - got %v", tt.want, got)
		}
	}
}
