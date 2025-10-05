package codility

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMinMovesUniqueCounts(t *testing.T) {
	cases := []struct {
		array []int
		want  int
	}{
		{[]int{1, 1, 3, 4, 4, 4}, 3},
		{[]int{1, 2, 2, 2, 5, 5, 5, 8}, 4},
		{[]int{1, 1, 1, 1, 3, 3, 4, 4, 4, 4, 4}, 5},
		{[]int{10, 10, 10}, 3},
	}

	countErr := 0
	for _, c := range cases {
		got := minMovesUniqueCounts(c.array)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("minMovesUniqueCounts test failed, want %v - got %v", c.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}

	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
