package leetcode_826

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxProfitAssignment(t *testing.T) {
	tests := []struct {
		difficulty []int
		profit     []int
		worker     []int
		want       int
	}{
		{[]int{7, 20, 68}, []int{26, 28, 57}, []int{71, 20, 71}, 142},
		{[]int{13, 37, 58}, []int{4, 90, 96}, []int{34, 73, 45}, 190},
		{[]int{2, 4, 6, 8, 10}, []int{10, 20, 30, 40, 50}, []int{4, 5, 6, 7}, 100},
		{[]int{85, 47, 57}, []int{24, 66, 99}, []int{40, 25, 25}, 0},
	}

	countErr := 0
	for _, tt := range tests {
		got := maxProfitAssignment(tt.difficulty, tt.profit, tt.worker)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("maxProfitAssignment test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
