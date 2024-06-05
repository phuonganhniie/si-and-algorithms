package leetcode_1002

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCommonChars(t *testing.T) {
	tests := []struct {
		words []string
		want  []string
	}{
		{[]string{"bella", "label", "roller"}, []string{"e", "l", "l"}},
		{[]string{"cool", "lock", "cook"}, []string{"c", "o"}},
	}

	countErr := 0
	for _, tt := range tests {
		got := commonChars(tt.words)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("commonChars test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
