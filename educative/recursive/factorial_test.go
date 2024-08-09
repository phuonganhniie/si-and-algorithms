package recursive

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		i    int
		want int
	}{
		{4, 24},
		{5, 120},
		{11, 39916800},
	}

	countErr := 0
	for _, tt := range tests {
		got := factorial(tt.i)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("factorial test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
