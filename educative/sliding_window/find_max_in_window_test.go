package slidingwindow

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{10, 6, 9, -3, 23, -1, 34, 56, 67, -1, -4, -8, -2, 9, 10, 34, 67}, 3, []int{10, 9, 23, 23, 34, 56, 67, 67, 67, -1, -2, 9, 10, 34, 67}},
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{[]int{1}, 1, []int{1}},
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 3, []int{10, 9, 8, 7, 6, 5, 4, 3}},
	}

	for _, tt := range tests {
		got := findMaxSlidingWindow(tt.nums, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("findMaxSlidingWindow test failed, want %v - got %v", tt.want, got)
		}
		fmt.Println("==================================")
	}
}
