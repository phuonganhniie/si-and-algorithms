package leetcode_3110

import (
	"fmt"
	"reflect"
	"testing"
)

func TestScoreOfString(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"hello", 13},
		{"zaz", 50},
	}

	countErr := 0
	for _, tt := range tests {
		got := scoreOfString(tt.s)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("scoreOfString test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
