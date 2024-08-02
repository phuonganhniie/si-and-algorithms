package leetcode_2678

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCountSeniors(t *testing.T) {
	tests := []struct {
		details []string
		want    int
	}{
		{[]string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}, 2},
		{[]string{"1313579440F2036", "2921522980M5644"}, 0},
	}

	countErr := 0
	for _, tt := range tests {
		got := countSeniors(tt.details)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("countSeniors test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
