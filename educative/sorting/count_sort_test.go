package sorting

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCountSort(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{[]int{2, 8, 3, 1, 4, 2, 4}, []int{1, 2, 2, 3, 4, 4, 8}},
		{[]int{23, 24, 22, 21, 26, 25, 27, 28, 21, 21}, []int{21, 21, 21, 22, 23, 24, 25, 26, 27, 28}},
	}

	countErr := 0
	for _, tt := range tests {
		got := countingSort(tt.arr)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("countingSort test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
