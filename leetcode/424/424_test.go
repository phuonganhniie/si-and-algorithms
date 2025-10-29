package leetcode_424

import (
	"fmt"
	"testing"
)

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{
			name: "Example 1: ABAB with 2 replacements",
			s:    "ABAB",
			k:    0,
			want: 1,
		},
		{
			name: "Example 2: AABABBA with 1 replacement",
			s:    "AABABBA",
			k:    1,
			want: 4,
		},
		{
			name: "All same characters",
			s:    "AAAA",
			k:    2,
			want: 4,
		},
		{
			name: "Single character string",
			s:    "A",
			k:    0,
			want: 1,
		},
		{
			name: "Two characters, no replacement allowed",
			s:    "AB",
			k:    0,
			want: 1,
		},
		{
			name: "Two characters, one replacement allowed",
			s:    "AB",
			k:    1,
			want: 2,
		},
		{
			name: "k equals string length",
			s:    "ABCD",
			k:    4,
			want: 4,
		},
		{
			name: "k greater than string length",
			s:    "ABC",
			k:    5,
			want: 3,
		},
		{
			name: "No replacements needed",
			s:    "AAAA",
			k:    0,
			want: 4,
		},
		{
			name: "Alternating characters with k=1",
			s:    "ABABABAB",
			k:    1,
			want: 3, // "ABA" or "BAB": 2 of one char, 1 of another → replace 1
		},
		{
			name: "Long sequence with one replacement",
			s:    "AAABAAAA",
			k:    1,
			want: 8,
		},
		{
			name: "Multiple different characters",
			s:    "ABCABC",
			k:    2,
			want: 4,
		},
		{
			name: "Best window at the beginning",
			s:    "AAABAAA",
			k:    1,
			want: 7,
		},
		{
			name: "Best window at the end",
			s:    "ABAAAA",
			k:    1,
			want: 6, // Entire string: 5 A's + 1 B, replace B → "AAAAAA"
		},
		{
			name: "Best window in the middle",
			s:    "BAAAAAB",
			k:    1,
			want: 6,
		},
		{
			name: "All different characters, k=0",
			s:    "ABCDEF",
			k:    0,
			want: 1,
		},
		{
			name: "All different characters, k=2",
			s:    "ABCDEF",
			k:    2,
			want: 3,
		},
		{
			name: "Complex pattern with k=2",
			s:    "AABABBABB",
			k:    2,
			want: 7,
		},
		{
			name: "Two character types only",
			s:    "AABBAABB",
			k:    2,
			want: 6, // Window "AABBAA": 4 A's, 2 B's → replace 2 B's
		},
		{
			name: "Long string with repeating pattern",
			s:    "ABBBAAABBBAAABBB",
			k:    2,
			want: 5, // With single-step shrink optimization: "BBBAA" = 3B, 2A → replace 2
		},
		{
			name: "String with three character types, k=1",
			s:    "AABCBBA",
			k:    1,
			want: 4,
		},
		{
			name: "Maximum replacement for short string",
			s:    "AAB",
			k:    1,
			want: 3,
		},
		{
			name: "Minimum k for diverse string",
			s:    "ABCABCABC",
			k:    1,
			want: 2, // Window "AB", "BC", or "CA": 1 of each type → replace 1 → 2 same
		},
		{
			name: "Large k value",
			s:    "ABCD",
			k:    10,
			want: 4,
		},
		{
			name: "Consecutive same chars with gaps",
			s:    "AAABBBAAABBB",
			k:    2,
			want: 5, // Window "AAABB" or "BBAAA": max 3 of one type + 2 replacements = 5
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := characterReplacement(tt.s, tt.k)
			if got != tt.want {
				t.Errorf("characterReplacement() = %v, want %v", got, tt.want)
				fmt.Printf("Input: s=%q, k=%d\n", tt.s, tt.k)
			} else {
				fmt.Printf("✓ %s: PASSED (result: %d)\n", tt.name, got)
			}
		})
	}
}

// Benchmark test to measure performance
func BenchmarkCharacterReplacement(b *testing.B) {
	s := "AABABBA"
	k := 1
	for i := 0; i < b.N; i++ {
		characterReplacement(s, k)
	}
}

// Benchmark with large input
func BenchmarkCharacterReplacementLarge(b *testing.B) {
	s := ""
	for i := 0; i < 10000; i++ {
		if i%3 == 0 {
			s += "A"
		} else if i%3 == 1 {
			s += "B"
		} else {
			s += "C"
		}
	}
	k := 100

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		characterReplacement(s, k)
	}
}
