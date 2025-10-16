package leetcode_2273

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveAnagrams(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  []string
	}{
		{
			name:  "Example 1: Multiple anagrams to remove",
			words: []string{"abba", "baba", "bbaa", "cd", "cd"},
			want:  []string{"abba", "cd"},
		},
		{
			name:  "Example 2: No anagrams",
			words: []string{"a", "b", "c", "d", "e"},
			want:  []string{"a", "b", "c", "d", "e"},
		},
		{
			name:  "All words are anagrams",
			words: []string{"abc", "bca", "cab", "acb"},
			want:  []string{"abc"},
		},
		{
			name:  "Single word",
			words: []string{"hello"},
			want:  []string{"hello"},
		},
		{
			name:  "Two identical words",
			words: []string{"test", "test"},
			want:  []string{"test"},
		},
		{
			name:  "Anagrams at the beginning",
			words: []string{"abc", "bca", "xyz", "foo"},
			want:  []string{"abc", "xyz", "foo"},
		},
		{
			name:  "Anagrams at the end",
			words: []string{"foo", "bar", "abc", "bca", "cab"},
			want:  []string{"foo", "bar", "abc"},
		},
		{
			name:  "Alternating anagrams and non-anagrams",
			words: []string{"abc", "bca", "xyz", "zyx", "hello"},
			want:  []string{"abc", "xyz", "hello"},
		},
		{
			name:  "Multiple groups of anagrams",
			words: []string{"dog", "god", "cat", "tac", "bird"},
			want:  []string{"dog", "cat", "bird"},
		},
		{
			name:  "Single character words with anagrams",
			words: []string{"a", "a", "b", "b", "c"},
			want:  []string{"a", "b", "c"},
		},
		{
			name:  "Long sequence of same anagrams",
			words: []string{"listen", "silent", "enlist", "inlets", "hello"},
			want:  []string{"listen", "hello"},
		},
		{
			name:  "Case with empty-like behavior - two words only",
			words: []string{"ab", "ba"},
			want:  []string{"ab"},
		},
		{
			name:  "Words with repeated letters",
			words: []string{"aabbcc", "abcabc", "cbacba", "xyz"},
			want:  []string{"aabbcc", "xyz"},
		},
		{
			name:  "Non-adjacent anagrams should remain",
			words: []string{"abc", "xyz", "bca"},
			want:  []string{"abc", "xyz", "bca"},
		},
		{
			name:  "Complex pattern with multiple removals",
			words: []string{"z", "z", "z", "gsw", "wsg", "gsw", "wsg"},
			want:  []string{"z", "gsw"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of input to avoid modification affecting the test output
			input := make([]string, len(tt.words))
			copy(input, tt.words)

			got := removeAnagrams(input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeAnagrams() = %v, want %v", got, tt.want)
				fmt.Printf("Input: %v\n", tt.words)
			} else {
				fmt.Printf("✓ %s: PASSED (result: %v)\n", tt.name, got)
			}
		})
	}
}

// Benchmark test to measure performance
func BenchmarkRemoveAnagrams(b *testing.B) {
	words := []string{"abba", "baba", "bbaa", "cd", "cd"}
	for i := 0; i < b.N; i++ {
		// Create a copy for each iteration to avoid side effects
		input := make([]string, len(words))
		copy(input, words)
		removeAnagrams(input)
	}
}

// Benchmark with large input
func BenchmarkRemoveAnagramsLarge(b *testing.B) {
	words := make([]string, 1000)
	for i := range words {
		if i%2 == 0 {
			words[i] = "abcdefgh"
		} else {
			words[i] = "hgfedcba"
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		input := make([]string, len(words))
		copy(input, words)
		removeAnagrams(input)
	}
}
