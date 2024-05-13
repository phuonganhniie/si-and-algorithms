package leetcode_20

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
	}

	for _, tt := range tests {
		got := IsValid(tt.s)
		if got != tt.want {
			t.Errorf("IsValid(%s) = %v, want %v", tt.s, got, tt.want)
		}
	}
}
