package leetcode_2037

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMinMovesToSeat(t *testing.T) {
	tests := []struct {
		seats    []int
		students []int
		want     int
	}{
		{[]int{5}, []int{1}, 4},
		{[]int{3, 1, 5}, []int{2, 7, 4}, 4},
		{[]int{4, 1, 5, 9}, []int{1, 3, 2, 6}, 7},
		{[]int{2, 2, 6, 6}, []int{1, 3, 2, 6}, 4},
	}

	countErr := 0
	for _, tt := range tests {
		got := minMovesToSeat(tt.seats, tt.students)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("minMovesToSeat test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
