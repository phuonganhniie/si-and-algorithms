package fastandslow

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCircularArrayLoop(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{-1, 2, 1, 2}, true},
		{[]int{-2, -3, 1, -3, -2}, true},
		{[]int{1, 3, -2, -4, 1}, true},
		{[]int{2, 1, -1, -2}, false},
		{[]int{5, 4, -2, -1, 3}, false},
		{[]int{1, 2, -3, 3, 4, 7, 1}, true},
		{[]int{3, 3, 1, -1, 2}, true},
	}

	countErr := 0
	for i, tt := range tests {
		got := circularArrayLoop(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("circularArrayLoop test [%d] failed, want %v - got %v", i+1, tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
