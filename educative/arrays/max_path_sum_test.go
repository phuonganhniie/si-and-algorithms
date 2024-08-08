package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	tests := []struct {
		arr1  []int
		size1 int
		arr2  []int
		size2 int
		want  int
	}{
		{[]int{12, 13, 18, 20, 22, 26, 70}, 7, []int{11, 15, 18, 19, 20, 26, 30, 31}, 8, 201},
		{[]int{0, 1, 2, 5}, 4, []int{2, 3, 8, 9, 10}, 5, 33},
	}

	countErr := 0
	for _, tt := range tests {
		got := maxPathSum(tt.arr1, tt.size1, tt.arr2, tt.size2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("maxPathSum2 test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
