package slidingwindow

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 1, 1, 1, 1, 1, 1, 1}, 11, 0},
		{[]int{2, 3, 1, 2, 4, 3, 1}, 6, 2},
		{[]int{2, 3, 1, 2, 4, 3}, 7, 2},
		{[]int{1, 4, 4}, 4, 1},
	}

	for _, tt := range tests {
		got := minSubArrayLen(tt.target, tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("minSubArrayLen test failed, want %v - got %v", tt.want, got)
		}
		fmt.Println("==================================")
	}
}
