package leetcode

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"racecar", "racecar"},
		{"babad", "aba"},
		{"cbbd", "bb"},
	}

	for _, tt := range tests {
		got := LongestPalindrome(tt.s)
		if got != tt.want {
			t.Errorf("LongestPalindrome(%s) = %v, want %v", tt.s, got, tt.want)
		}
	}
}
