package leetcode_1071

import "testing"

func TestGcdOfStrings(t *testing.T) {
	cases := []struct {
		s    string
		t    string
		want string
	}{
		{"LEET", "CODE", ""},
		{"ABCABC", "ABC", "ABC"},
		{"ABABAB", "ABAB", "AB"},
	}

	for _, c := range cases {
		got := gcdOfStrings(c.s, c.t)
		if got != c.want {
			t.Errorf("gcdOfStrings test failed, want %v - got %v", c.want, got)
		}
	}
}
