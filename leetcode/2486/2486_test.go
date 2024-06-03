package leetcode_2486

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAppendCharacters(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want int
	}{
		{"coaching", "coding", 4},
		{"abcde", "a", 0},
		{"z", "abcde", 5},
	}

	countErr := 0
	for _, tt := range tests {
		got := appendCharacters(tt.s, tt.t)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("appendCharacters test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
