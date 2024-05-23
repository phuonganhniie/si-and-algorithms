package leetcode_392

import "testing"

func TestIsSubsequence(t *testing.T) {
	cases := []struct {
		s    string
		t    string
		want bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
		{"ace", "abcde", true},
		{"aec", "abcde", false},
	}

	for _, c := range cases {
		got := isSubsequence(c.s, c.t)
		if got != c.want {
			t.Errorf("IsPalindrome test fail, want %v - got %v", c.want, got)
		}
	}
}
