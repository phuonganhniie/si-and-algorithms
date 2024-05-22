package twopointers

import (
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"kaYak", true},
		{"hello", false},
		{"RaCEACAR", false},
	}

	for _, tt := range tests {
		got := isPalindrome(tt.s)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("isPalindrome test failed, want %v - got %v", tt.want, got)
		}
	}
}
