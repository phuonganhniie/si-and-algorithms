package arrays

import (
	"reflect"
	"testing"
)

func TestWaveForm(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{8, 1, 2, 3, 4, 5, 6, 4, 2}, []int{8, 1, 3, 2, 5, 4, 6, 2, 4}},
		{[]int{7, 8, 9, 10}, []int{8, 7, 10, 9}},
		{[]int{3, 4, 5, 6, 7, 8, 9}, []int{4, 3, 6, 5, 8, 7, 9}},
		{[]int{7, 3, 9, 10, 3}, []int{7, 3, 10, 3, 9}},
	}

	for _, tt := range tests {
		got := WaveArray(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("WaveArray test failed, want %v - got %v", tt.want, got)
		}
	}
}
