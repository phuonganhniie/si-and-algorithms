package twopointers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"tcaac", true},
		{"A", true},
		{"eeccccbebaeeabebccceea", false},
		{"madame", true},
		{"dead", true},
		{"tebbem", false},
		{"abca", true},
	}

	countErr := 0
	for _, tt := range tests {
		got := validPalindrome(tt.s)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("validPalindrome test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
