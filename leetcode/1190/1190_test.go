package leetcode_1190

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverseParentheses(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"(abcd)", "dcba"},
		{"(u(love)i)", "iloveu"},
		{"(ed(et(oc))el)", "leetcode"},
	}

	countErr := 0
	for _, tt := range tests {
		got := reverseParentheses(tt.s)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("reverseParentheses test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
