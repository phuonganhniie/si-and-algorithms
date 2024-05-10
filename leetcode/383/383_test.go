package leetcode_383

import (
	"testing"
)

func TestCanConstruct(t *testing.T) {
	tests := []struct {
		ransomNote string
		magazine   string
		want       bool
	}{
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},
		{"ab", "aaba", true},
	}

	for _, tt := range tests {
		got := CanConstruct(tt.ransomNote, tt.magazine)
		if got != tt.want {
			t.Errorf("CanConstruct test failed, want %v - got %v", tt.want, got)
		}
	}
}
