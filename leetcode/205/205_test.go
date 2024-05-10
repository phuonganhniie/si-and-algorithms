package leetcode_205

import "testing"

func TestIsIsomorphic(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{"\u0000\u0000", "\u0000c", false},
		{"egg", "add", true},
		{"foo", "bar", false},
		{"paper", "title", true},
	}

	for _, tt := range tests {
		got := IsIsomorphic(tt.s, tt.t)
		if got != tt.want {
			t.Errorf("IsIsomorphic test failed, want %v - got %v", tt.want, got)
		}
	}
}
