package leetcode_846

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsNStraightHand(t *testing.T) {
	tests := []struct {
		hand      []int
		groupSize int
		want      bool
	}{
		{[]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3, true},
		{[]int{1, 2, 3, 4, 5}, 4, false},
	}

	countErr := 0
	for _, tt := range tests {
		got := isNStraightHand(tt.hand, tt.groupSize)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("isNStraightHand test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
