package stack

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/phuonganhniie/educative/helper"
)

func TestBottomInsert(t *testing.T) {
	tests := []struct {
		stack   *helper.Stack
		element int
		want    *helper.Stack
	}{
		{&helper.Stack{Arr: []int{1, 2, 3}}, 5, &helper.Stack{Arr: []int{5, 1, 2, 3}}},
		{&helper.Stack{Arr: []int{8, 5, 1, 2, 3}}, 0, &helper.Stack{Arr: []int{0, 8, 5, 1, 2, 3}}},
	}

	countErr := 0
	for _, tt := range tests {
		got := bottomInsert(tt.stack, tt.element)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("bottomInsert test failed, want: %v - got: %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
