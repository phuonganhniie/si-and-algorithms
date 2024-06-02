package leetcode_334

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIncreasingTriplet(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{5, 4, 3, 2, 1}, false},
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{2, 1, 5, 0, 4, 6}, true},
		{[]int{2, 1}, false},
	}

	countErr := 0
	for _, tt := range tests {
		got := increasingTriplet(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("increasingTriplet test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
