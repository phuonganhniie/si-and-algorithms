package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxMinArray(t *testing.T) {
	tests := []struct {
		arr  []int
		size int
		want []int
	}{
		{[]int{4, 8, 16, 20, 24}, 5, []int{24, 4, 20, 8, 16}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, []int{10, 1, 9, 2, 8, 3, 7, 4, 6, 5}},
	}

	countErr := 0
	for _, tt := range tests {
		got := maxMinArr(tt.arr, tt.size)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("maxMinArr test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
