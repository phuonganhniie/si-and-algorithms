package leetcode_1768

import (
	"reflect"
	"testing"
)

func TestMergeAlternately(t *testing.T) {
	tests := []struct {
		word1 string
		word2 string
		want  string
	}{
		{"abc", "pqr", "apbqcr"},
		{"ab", "pqrs", "apbqrs"},
	}

	for _, tt := range tests {
		got := MergeAlternately(tt.word1, tt.word2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("MergeAlternately test failed, want %v - got %v", tt.want, got)
		}
	}
}
