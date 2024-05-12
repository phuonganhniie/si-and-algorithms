package leetcode_2373

import (
	"reflect"
	"testing"
)

func TestLargestLocal(t *testing.T) {
	cases := []struct {
		grid [][]int
		want [][]int
	}{
		{[][]int{{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2}}, [][]int{{9, 9}, {8, 6}}},
		{[][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 2, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}}, [][]int{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}}},
	}

	for _, c := range cases {
		got := LargestLocal(c.grid)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("LargestLocal test failed, want %d - got %d", c.want, got)
		}
	}
}
