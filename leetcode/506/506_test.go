package leetcode_506

import (
	"reflect"
	"testing"
)

func TestFindRelativeRanks(t *testing.T) {
	cases := []struct {
		score []int
		want  []string
	}{
		{[]int{10, 3, 8, 9, 4}, []string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"}},
		{[]int{5, 4, 3, 2, 1}, []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"}},
	}

	for _, c := range cases {
		got := findRelativeRanks(c.score)
		if !reflect.DeepEqual(c.want, got) {
			t.Errorf("findRelativeRanks test failed, want %v - got %v", c.want, got)
		}
	}
}
