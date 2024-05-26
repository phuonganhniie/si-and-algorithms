package leetcode_1456

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxVowels(t *testing.T) {
	tests := []struct {
		s    string
		k    int
		want int
	}{
		{"ibpbhixfiouhdljnjfflpapptrxgcomvnb", 33, 7},
		{"abciiidef", 3, 3},
		{"aeiou", 2, 2},
		{"leetcode", 3, 2},
	}

	for _, tt := range tests {
		got := maxVowels(tt.s, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("maxVowels test failed, want %v - got %v", tt.want, got)
		}
		fmt.Println("==================================")
	}
}
