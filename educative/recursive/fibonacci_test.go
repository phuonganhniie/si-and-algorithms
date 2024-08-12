package recursive

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{4, 3},
		{6, 8},
		{11, 89},
		{27, 196418},
	}

	countErr := 0
	for _, tt := range tests {
		got := fibonacci(tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("fibonacci test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
