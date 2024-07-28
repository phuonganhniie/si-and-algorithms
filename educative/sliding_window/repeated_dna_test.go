package slidingwindow

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindRepeatedSequences(t *testing.T) {
	tests := []struct {
		dna  string
		k    int
		want []string
	}{
		{"AAAAACCCCCAAAAACCCCCC", 8, []string{"AAAAACCC", "AAAACCCC", "AAACCCCC"}},
		{"GGGGGGGGGGGGGGGGGGGGGGGGG", 9, []string{"GGGGGGGGG"}},
		{"TTTTTCCCCCCCTTTTTTCCCCCCCTTTTTTT", 10, []string{"CCCCCCCTTT", "CCCCCCTTTT", "CCCCCTTTTT", "CCCCTTTTTT", "TCCCCCCCTT", "TTCCCCCCCT", "TTTCCCCCCC", "TTTTCCCCCC", "TTTTTCCCCC"}},
		{"AAAAAACCCCCCCAAAAAAAACCCCCCCTG", 10, []string{"AAAAAACCCC", "AAAAACCCCC", "AAAACCCCCC", "AAACCCCCCC"}},
		{"ATATATATATATATAT", 6, []string{"ATATAT", "TATATA"}},
	}

	countErr := 0
	for _, tt := range tests {
		rs := findRepeatedSequences(tt.dna, tt.k)
		got := rs.SetToArrayStr()
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("findRepeatedSequences test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}

	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
