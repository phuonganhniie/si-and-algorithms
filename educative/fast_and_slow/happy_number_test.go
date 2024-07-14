package fastandslow

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsHappy(t *testing.T) {
	tests := []struct {
		num  int
		want bool
	}{
		{2147483646, false},
		{1, true},
		{19, true},
		{8, false},
		{7, true},
	}

	countErr := 0
	for _, tt := range tests {
		got := isHappy(tt.num)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("isHappy test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
