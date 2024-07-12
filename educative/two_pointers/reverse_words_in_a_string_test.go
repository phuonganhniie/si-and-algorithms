package twopointers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		sentence string
		want     string
	}{
		{"We love Go", "Go love We"},
		{"1234 abc XYZ", "XYZ abc 1234"},
		{"You are amazing", "amazing are You"},
		{"Hello     World", "World Hello"},
		{"Greeting123", "Greeting123"},
	}

	countErr := 0
	for _, tt := range tests {
		got := reverseWords2(tt.sentence)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("reverseWords test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
