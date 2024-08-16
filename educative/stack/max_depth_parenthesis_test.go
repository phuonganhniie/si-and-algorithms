package stack

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxDepthParenthesis(t *testing.T) {
	tests := []struct {
		exp  string
		size int
		want int
	}{
		{"(()(())())", 8, 3},
		{"()", 2, 1},
		{"((()))", 6, 3},
		{"((((A)))((((BBB()))))()()()())", 30, 6},
	}

	countErr := 0
	for _, tt := range tests {
		got := maxDepthParenthesis(tt.exp, tt.size)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("maxDepthParenthesis test failed, want: %v - got: %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
