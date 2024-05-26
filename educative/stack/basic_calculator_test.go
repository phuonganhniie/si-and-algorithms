package stack

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCalculator(t *testing.T) {
	tests := []struct {
		expression string
		want       int
	}{
		{"12 - (6 + 2) + 5", 9},
		{"(8 + 100) + (13 - 8 - (2 + 1))", 110},
		{"40 - 25 - 5", 10},
		{"(27 + (7 + 5) - 3) + (6 + 10)", 52},
		{"34 + 45 + (23324 - 3454)", 19949},
	}

	for _, tt := range tests {
		got := calculator(tt.expression)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("calculator test failed, want %v - got %v", tt.want, got)
		}
		fmt.Println("==================================")
	}
}
