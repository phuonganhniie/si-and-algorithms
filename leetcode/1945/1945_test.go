package leetcode_1945

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetLucky(t *testing.T) {
	tests := []struct {
		s    string
		k    int
		want int
	}{
		{"iaozzbyqzwbpurzze", 2, 5},
		{"leetcode", 2, 6},
		{"iiii", 1, 36},
		{"zbax", 2, 8},
	}

	countErr := 0

	for _, tt := range tests {
		got := getLucky(tt.s, tt.k)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("getLucky test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}

	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
