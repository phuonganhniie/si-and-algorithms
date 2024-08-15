package stack

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/phuonganhniie/educative/helper"
)

func TestReverseStack(t *testing.T) {
	tests := []struct {
		stack *helper.Stack
		want  *helper.Stack
	}{
		{&helper.Stack{Arr: []int{1, 2, 3, 4, 5}}, &helper.Stack{Arr: []int{5, 4, 3, 2, 1}}},
	}

	countErr := 0
	for _, tt := range tests {
		got := reverseStack(tt.stack)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("reverseStack test failed, want: %v - got: %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
