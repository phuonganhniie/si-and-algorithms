package leetcode_1518

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNumWaterBottles(t *testing.T) {
	tests := []struct {
		numBottles  int
		numExchange int
		want        int
	}{
		{9, 3, 13},
		{15, 4, 19},
	}

	countErr := 0
	for _, tt := range tests {
		got := numWaterBottles2(tt.numBottles, tt.numExchange)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("numWaterBottles test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
