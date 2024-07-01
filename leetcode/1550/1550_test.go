package leetcode_1550

import (
	"fmt"
	"reflect"
	"testing"
)

func TestThreeConsecutiveOdds(t *testing.T) {
	tests := []struct {
		arr  []int
		want bool
	}{
		{[]int{2, 6, 4, 1}, false},
		{[]int{1, 2, 34, 3, 4, 5, 7, 23, 12}, true},
	}

	countErr := 0
	for _, tt := range tests {
		got := threeConsecutiveOdds(tt.arr)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("threeConsecutiveOdds test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
