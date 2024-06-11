package leetcode_1122

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRelativeSortArray(t *testing.T) {
	tests := []struct {
		arr1 []int
		arr2 []int
		want []int
	}{
		{[]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}, []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19}},
		{[]int{28, 6, 22, 8, 44, 17}, []int{22, 28, 8, 6}, []int{22, 28, 8, 6, 17, 44}},
	}

	countErr := 0
	for _, tt := range tests {
		got := relativeSortArrayCountingSort(tt.arr1, tt.arr2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("relativeSortArray test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
