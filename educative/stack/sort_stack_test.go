package stack

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/phuonganhniie/educative/helper"
)

func TestSortStack(t *testing.T) {
	tests := []struct {
		stack *helper.Stack
		want  *helper.Stack
	}{
		{&helper.Stack{Arr: []int{1, 2, 4, 9}}, &helper.Stack{Arr: []int{1, 2, 4, 9}}},
		{&helper.Stack{Arr: []int{4, 8, 2, 7, 1}}, &helper.Stack{Arr: []int{1, 2, 4, 7, 8}}},
	}

	countErr := 0
	for _, tt := range tests {
		got := sortStack(tt.stack)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("sortStack test failed, want: %v - got: %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
