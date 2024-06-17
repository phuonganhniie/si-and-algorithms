package leetcode_633

import (
	"fmt"
	"reflect"
	"testing"
)

func TestJudgeSquareSum(t *testing.T) {
	tests := []struct {
		c    int
		want bool
	}{
		{5, true},
		{3, false},
	}

	countErr := 0
	for _, tt := range tests {
		got := judgeSquareSum(tt.c)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("judgeSquareSum test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
