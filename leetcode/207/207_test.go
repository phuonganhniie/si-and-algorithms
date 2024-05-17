package leetcode_207

import (
	"reflect"
	"testing"
)

func TestCanFinish(t *testing.T) {
	cases := []struct {
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{4, [][]int{{1, 0}, {2, 1}, {3, 2}}, true},
		{2, [][]int{{1, 0}}, true},
		{2, [][]int{{1, 0}, {0, 1}}, false},
	}

	for _, c := range cases {
		got := CanFinish(c.numCourses, c.prerequisites)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("CanFinish test failed, want %v - got %v", c.want, got)
		}
	}
}
