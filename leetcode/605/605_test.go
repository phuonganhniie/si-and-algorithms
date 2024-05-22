package leetcode_605

import (
	"reflect"
	"testing"
)

func TestCanPlaceFlowers(t *testing.T) {
	tests := []struct {
		flowerbed []int
		n         int
		want      bool
	}{
		{[]int{0, 0, 1, 0, 0}, 1, true},
		{[]int{0}, 1, true},
		{[]int{1, 0, 0, 0, 1}, 1, true},
		{[]int{1, 0, 0, 0, 1}, 2, false},
	}

	for _, tt := range tests {
		got := CanPlaceFlowers(tt.flowerbed, tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("CanPlaceFlowers test failed, want %v - got %v", tt.want, got)
		}
	}
}
