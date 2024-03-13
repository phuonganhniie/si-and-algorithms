package leetcode_2485

import "testing"

func TestPivotIn(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{8, 6},
		{1, 1},
		{4, -1},
	}

	for _, tt := range tests {
		got := PivotIntegerBS(tt.input)
		if got != tt.want {
			t.Errorf("pivotInteger(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}
