package leetcode_1431

import (
	"reflect"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	tests := []struct {
		candies      []int
		extraCandies int
		want         []bool
	}{
		{[]int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},
		{[]int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false}},
		{[]int{12, 1, 12}, 10, []bool{true, false, true}},
	}

	for _, tt := range tests {
		got := KidsWithCandiesOptimized(tt.candies, tt.extraCandies)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("KidsWithCandies test failed, want %v - got %v", tt.want, got)
		}
	}
}
