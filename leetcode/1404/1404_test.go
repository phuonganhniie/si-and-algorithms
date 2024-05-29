package leetcode_1404

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNumSteps(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"1111110011101010110011100100101110010100101110111010111110110010", 89},
		{"1101", 6},
		{"10", 1},
		{"1", 0},
	}

	countErr := 0
	for _, tt := range tests {
		got := numSteps(tt.s)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("numSteps test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
