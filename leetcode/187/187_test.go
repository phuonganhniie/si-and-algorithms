package leetcode_187

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindRepeatedDnaSequences(t *testing.T) {
	tests := []struct {
		s    string
		want []string
	}{
		{"AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT", []string{"AAAAACCCCC", "CCCCCAAAAA"}},
		{"AAAAAAAAAAAAA", []string{"AAAAAAAAAA"}},
	}

	countErr := 0
	for _, tt := range tests {
		got := findRepeatedDnaSequences(tt.s)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("findRepeatedDnaSequences test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
