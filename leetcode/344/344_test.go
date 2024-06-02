package leetcode_344

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		s    []byte
		want string
	}{
		{[]byte{'h', 'e', 'l', 'l', 'o'}, "olleh"},
		{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}, "hannaH"},
	}

	countErr := 0
	for _, tt := range tests {
		got := reverseString(tt.s)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("reverseString test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
