package leetcode_165

import (
	"testing"
)

func TestCompareVersion(t *testing.T) {
	cases := []struct {
		version1 string
		version2 string
		want     int
	}{
		{"1.01", "1.001", 0},
		{"1.0", "1.0.0", 0},
		{"0.1", "1.1", -1},
	}

	for _, c := range cases {
		got := CompareVersion(c.version1, c.version2)
		if got != c.want {
			t.Errorf("CompareVersion(%s, %s) = %d, want %d got %d", c.version1, c.version2, got, c.want, got)
		}
	}
}
