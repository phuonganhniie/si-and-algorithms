package leetcode_1823

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindTheWinner(t *testing.T) {
	tests := []struct {
		n    int
		k    int
		want int
	}{
		{5, 2, 3},
		{6, 5, 1},
	}

	countErr := 0
	for _, tt := range tests {
		got := findTheWinner(tt.n, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("findTheWinner test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
