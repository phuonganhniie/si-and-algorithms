package leetcode_648

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReplaceWords(t *testing.T) {
	tests := []struct {
		dictionary []string
		sentence   string
		want       string
	}{
		{[]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery", "the cat was rat by the bat"},
		{[]string{"a", "b", "c"}, "aadsfasf absbs bbab cadsfafs", "a a b c"},
	}

	countErr := 0
	for _, tt := range tests {
		got := replaceWords(tt.dictionary, tt.sentence)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("replaceWords test failed, want %v - got %v", tt.want, got)
			fmt.Println("==================================")
			countErr++
		}
	}
	if countErr == 0 {
		fmt.Println("Chúp mừng pà nkaaaa")
	}
}
