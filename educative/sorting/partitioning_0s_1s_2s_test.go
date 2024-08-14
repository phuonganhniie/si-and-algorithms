package sorting

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPartition012(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{[]int{0, 0, 2, 0, 2, 1, 0, 1}, []int{0, 0, 0, 0, 1, 1, 2, 2}},
		{[]int{0, 0, 2, 0, 2, 1, 0, 1}, []int{0, 0, 0, 0, 1, 1, 2, 2}},
		{[]int{0, 1, 2, 0, 1, 2}, []int{0, 0, 1, 1, 2, 2}},
	}

	countErr := 0
	for _, tt := range tests {
		got := partition012Optimize(tt.arr)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("partition012Optimize test failed, want: %v - got: %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
