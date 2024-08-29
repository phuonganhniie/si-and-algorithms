package twopointers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRainWater(t *testing.T) {
	tests := []struct {
		heights []int
		want    int
	}{
		{[]int{1, 3, 0, 1, 3}, 5},
		{[]int{1, 0, 3, 0, 1, 2}, 4},
		{[]int{4, 2, 0, 3, 1, 5}, 10},
		{[]int{0, 3, 0, 2, 1, 0, 1, 4, 2, 1, 2, 0}, 12},
		{[]int{1, 2, 2, 2, 2, 2}, 0},
		{[]int{4}, 0},
	}

	countErr := 0
	for _, tt := range tests {
		got := rainWater(tt.heights)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("rainWater test failed, want: %v - got: %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}

	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}

func BenchmarkRainWater(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rainWater([]int{0, 3, 0, 2, 1, 0, 1, 4, 2, 1, 2, 0})
	}
}

func BenchmarkRainWater2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rainWater2([]int{0, 3, 0, 2, 1, 0, 1, 4, 2, 1, 2, 0})
	}
}
