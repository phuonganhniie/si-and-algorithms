package leetcode_125

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		s    string
		want bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{"  ", true},
	}

	for _, c := range cases {
		got := IsPalindrome(c.s)
		if got != c.want {
			t.Errorf("IsPalindrome test fail, want %v - got %v", c.want, got)
		}
	}
}
