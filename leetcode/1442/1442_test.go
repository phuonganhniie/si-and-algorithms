package leetcode_1442

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCountTriplets(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{[]int{2, 3, 1, 6, 7}, 4},
		{[]int{1, 1, 1, 1, 1}, 10},
	}

	countErr := 0
	for _, tt := range tests {
		got := countTriplets(tt.arr)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("countTriplets test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
