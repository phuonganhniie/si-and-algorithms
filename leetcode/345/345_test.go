package leetcode_345

import (
	"reflect"
	"testing"
)

func TestReverseVowels(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"hello", "holle"},
		{"leetcode", "leotcede"},
		{"aA", "Aa"},
	}

	for _, tt := range tests {
		got := ReverseVowels(tt.s)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("ReverseVowels test failed, want %v - got %v", tt.want, got)
		}
	}
}
