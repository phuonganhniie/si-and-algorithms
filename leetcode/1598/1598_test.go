package leetcode_1598

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMinOperations(t *testing.T) {
	tests := []struct {
		logs []string
		want int
	}{
		{[]string{"d1/", "d2/", "../", "d21/", "./"}, 2},
		{[]string{"d1/", "d2/", "./", "d3/", "../", "d31/"}, 3},
		{[]string{"d1/", "../", "../", "../"}, 0},
	}

	countErr := 0
	for _, tt := range tests {
		got := minOperations(tt.logs)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("minOperations test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
