package leetcode_1482

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMinDays(t *testing.T) {
	tests := []struct {
		bloomDay []int
		bouquets int
		flowers  int
		want     int
	}{
		{[]int{7, 7, 7, 7, 12, 7, 7}, 2, 3, 12},
		{[]int{1, 10, 3, 10, 2}, 3, 1, 3},
		{[]int{1, 10, 3, 10, 2}, 3, 2, -1},
	}

	countErr := 0
	for _, tt := range tests {
		got := minDays(tt.bloomDay, tt.bouquets, tt.flowers)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("minDays test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
