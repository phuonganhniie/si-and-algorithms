package leetcode_03

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}

	for _, tt := range tests {
		got := LengthOfLongestSubstring(tt.s)
		if got != tt.want {
			t.Errorf("LengthOfLongestSubstring(%s) = %v, want %v", tt.s, got, tt.want)
		}
	}
}
