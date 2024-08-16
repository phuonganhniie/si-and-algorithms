package stack

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsBalancedParenthesis(t *testing.T) {
	tests := []struct {
		exp  string
		want bool
	}{
		{"{()}[", false},
		{"{() ({})}", true},
		{"{[({})]}", true},
		{"((}))", false},
		{"{[}]", false},
	}

	countErr := 0
	for _, tt := range tests {
		got := isBalancedParenthesis(tt.exp)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("isBalancedParenthesis test failed, want: %v - got: %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
