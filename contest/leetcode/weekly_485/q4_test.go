package weekly485

import "testing"

func Test_smallestString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "Example 1",
			s:    "aaccb",
			want: "aacb",
		},
		{
			name: "Example 2",
			s:    "z",
			want: "z",
		},
		{
			name: "All same characters",
			s:    "aaaa",
			want: "a",
		},
		{
			name: "Descending order",
			s:    "dcba",
			want: "dcba",
		},
		{
			name: "Already sorted",
			s:    "abcd",
			want: "abcd",
		},
		{
			name: "Complex case",
			s:    "bcabc",
			want: "abc",
		},
		{
			name: "Another complex case",
			s:    "cbacdcbc",
			want: "acdb",
		},
		{
			name: "Another complex case 2",
			s:    "abcab",
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestString(tt.s); got != tt.want {
				t.Errorf("smallestString() = %v, want %v", got, tt.want)
			}
		})
	}
}
