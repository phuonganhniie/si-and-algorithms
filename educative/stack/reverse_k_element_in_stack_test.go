package stack

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/phuonganhniie/educative/helper"
)

func TestReverseKElementInStack(t *testing.T) {
	tests := []struct {
		stack *helper.Stack
		size  int
		want  *helper.Stack
	}{
		{&helper.Stack{Arr: []int{1, 2, 3, 4, 5}}, 4, &helper.Stack{Arr: []int{1, 5, 4, 3, 2}}},
	}

	countErr := 0
	for _, tt := range tests {
		got := reverseKElementInStack(tt.stack, tt.size)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("reverseKElementInStack test failed, want: %v - got: %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
