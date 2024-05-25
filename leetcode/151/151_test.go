package leetcode_151

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "world hello"},
		{"a good   example", "example good a"},
	}

	for _, tt := range tests {
		got := reverseWords(tt.s)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("reverseWords test failed, want %v - got %v", tt.want, got)
		}
		fmt.Println("==================================")
	}
}
