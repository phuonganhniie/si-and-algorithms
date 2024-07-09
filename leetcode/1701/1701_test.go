package leetcode_1701

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAverageWaitingTime(t *testing.T) {
	tests := []struct {
		customers [][]int
		want      float64
	}{
		{[][]int{{1, 2}, {2, 5}, {4, 3}}, 5.00000},
		{[][]int{{5, 2}, {5, 4}, {10, 3}, {20, 1}}, 3.25000},
	}

	countErr := 0
	for _, tt := range tests {
		got := averageWaitingTime2(tt.customers)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("averageWaitingTime test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
