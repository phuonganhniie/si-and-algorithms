package fastandslow

import (
	"reflect"
	"testing"
)

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 3, 6, 2, 7, 3, 5, 4}, 3},
		{[]int{3, 3, 3, 3, 3}, 3},
		{[]int{3, 4, 4, 4, 2}, 4},
		{[]int{1, 1}, 1},
		{[]int{1, 3, 4, 2, 2}, 2},
		{[]int{1, 2, 2}, 2},
	}

	for _, tt := range tests {
		got := findDuplicate(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("findDuplicate test failed, want %v - got %v", tt.want, got)
		}
	}
}
