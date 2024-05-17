package leetcode_412

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		n    int
		want []string
	}{
		{3, []string{"1", "2", "Fizz"}},
		{5, []string{"1", "2", "Fizz", "4", "Buzz"}},
		{15, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
	}

	for _, c := range cases {
		got := FizzBuzz(c.n)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("FizzBuzz test fail, want %v - got %v", c.want, got)
		}
	}
}
