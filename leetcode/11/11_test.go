package leetcode_11

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
	}

	for _, tt := range tests {
		got := maxArea(tt.height)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("maxArea test failed, want %v - got %v", tt.want, got)
		}
		fmt.Println("==================================")
	}
}
