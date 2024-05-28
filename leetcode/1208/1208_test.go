package leetcode_1208

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEqualSubstringBinarySearch(t *testing.T) {
	tests := []struct {
		s       string
		t       string
		maxCost int
		want    int
	}{
		{"abcd", "cdef", 1, 0},
		{"abcd", "bcdf", 3, 3},
		{"abcd", "cdef", 3, 1},
		{"abcd", "acde", 0, 1},
	}

	countErr := 0
	for _, tt := range tests {
		got := equalSubstringBinarySearch(tt.s, tt.t, tt.maxCost)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("equalSubstringBinarySearch test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}

func TestEqualSubstringSlidingWindow(t *testing.T) {
	tests := []struct {
		s       string
		t       string
		maxCost int
		want    int
	}{
		{"abcd", "bcdf", 3, 3},
		{"abcd", "cdef", 3, 1},
		{"abcd", "acde", 0, 1},
	}

	countErr := 0
	for _, tt := range tests {
		got := equalSubstringSlidingWindow(tt.s, tt.t, tt.maxCost)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("equalSubstringSlidingWindow test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
