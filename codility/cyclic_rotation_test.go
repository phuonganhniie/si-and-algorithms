package codility

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCyclicRotation(t *testing.T) {
	tests := []struct {
		A    []int
		K    int
		want []int
	}{
		{[]int{3, 8, 9, 7, 6}, 3, []int{9, 7, 6, 3, 8}},
		{[]int{1, 2, 3, 4}, 4, []int{1, 2, 3, 4}},
		{[]int{0, 0, 0}, 1, []int{0, 0, 0}},
	}

	countErr := 0
	for _, tt := range tests {
		got := cyclicRotation(tt.A, tt.K)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("cyclicRotation test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
