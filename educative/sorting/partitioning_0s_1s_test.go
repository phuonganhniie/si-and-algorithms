package sorting

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPartition01(t *testing.T) {
	tests := []struct {
		arr   []int
		size  int
		want1 []int
		want2 int
	}{
		{[]int{1, 0, 1, 0, 1, 0, 1}, 7, []int{0, 0, 0, 1, 1, 1, 1}, 2},
		{[]int{0, 1, 1, 0, 1, 0, 1, 1, 0, 0, 0, 1}, 12, []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1}, 3},
		{[]int{0, 1, 1, 1}, 4, []int{0, 1, 1, 1}, 0},
	}

	countErr := 0
	for _, tt := range tests {
		got1, got2 := partition01Optimize(tt.arr, tt.size)
		if !reflect.DeepEqual(got1, tt.want1) || !reflect.DeepEqual(got2, tt.want2) {
			t.Errorf("partition01 test failed, want1: %v - got1: %v ----- want2: %v - got2: %v", tt.want1, got1, tt.want2, got2)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
