package stack

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/phuonganhniie/educative/helper"
)

func TestSortedInsert(t *testing.T) {
	tests := []struct {
		arr     *helper.Stack
		element int
		want    *helper.Stack
	}{
		{&helper.Stack{Arr: []int{1, 2, 4, 9}}, 3, &helper.Stack{Arr: []int{1, 2, 3, 4, 9}}},
	}

	countErr := 0
	for _, tt := range tests {
		got := sortedInsert(tt.arr, tt.element)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("sortedInsert test failed, want: %v - got: %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
