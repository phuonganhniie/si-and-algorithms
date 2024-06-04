package leetcode_409

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abccccdd", 7},
		{"a", 1},
	}

	countErr := 0
	for _, tt := range tests {
		got := longestPalindrome(tt.s)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("longestPalindrome test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
