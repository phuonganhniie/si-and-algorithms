package leetcode_861

import (
	"reflect"
	"testing"
)

func TestMatrixScore(t *testing.T) {
	cases := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}}, 39},
		{[][]int{{0}}, 1},
	}

	for _, c := range cases {
		got := MatrixScore(c.grid)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("MatrixScore test failed, want %d - got %d", c.want, got)
		}
	}
}
