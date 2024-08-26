package twopointers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},

		{[]int{5, 3, -1, 6, 4}, []int{-72, -120, 360, -60, -90}},
		{[]int{0, -1, 2, -3, 4, -2}, []int{-48, 0, 0, 0, 0, 0}},
	}

	countErr := 0
	for _, tt := range tests {
		got := productExceptSelfEducative(tt.arr)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("productExceptSelf test failed, want: %v - got: %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
