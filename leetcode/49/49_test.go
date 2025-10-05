package leetcode_49

import (
	"fmt"
	"strings"
	"testing"

	"slices"
)

func normalizeGroups(groups [][]string) [][]string {
	normalized := make([][]string, 0, len(groups))
	for _, g := range groups {
		cp := append([]string(nil), g...)
		slices.Sort(cp)
		normalized = append(normalized, cp)
	}
	slices.SortFunc(normalized, func(a, b []string) int {
		aj := strings.Join(a, ",")
		bj := strings.Join(b, ",")
		switch {
		case aj < bj:
			return -1
		case aj > bj:
			return 1
		default:
			return 0
		}
	})
	return normalized
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  [][]string
	}{
		{
			name:  "Example: basic anagrams",
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want:  [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			name:  "Empty input",
			input: []string{},
			want:  [][]string{},
		},
		{
			name:  "Single empty string",
			input: []string{""},
			want:  [][]string{{""}},
		},
		{
			name:  "Duplicates included",
			input: []string{"abc", "bca", "cab", "abc"},
			want:  [][]string{{"abc", "abc", "bca", "cab"}},
		},
		{
			name:  "Different lengths",
			input: []string{"a", "ab", "ba", "abc", "cab", "bca", "z"},
			want:  [][]string{{"a"}, {"z"}, {"ab", "ba"}, {"abc", "bca", "cab"}},
		},
		{
			name:  "All singletons (no anagrams)",
			input: []string{"abc", "def", "ghi"},
			want:  [][]string{{"abc"}, {"def"}, {"ghi"}},
		},
		{
			name:  "Repeated same word",
			input: []string{"aaaa", "aaaa", "aaaa"},
			want:  [][]string{{"aaaa", "aaaa", "aaaa"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.input)

			nGot := normalizeGroups(got)
			nWant := normalizeGroups(tt.want)

			eq := len(nGot) == len(nWant)
			if eq {
				for i := range nGot {
					if !slices.Equal(nGot[i], nWant[i]) {
						eq = false
						break
					}
				}
			}

			if !eq {
				t.Errorf("groupAnagrams() mismatch.\ninput: %v\n got: %v\nwant: %v", tt.input, nGot, nWant)
			} else {
				fmt.Printf("✓ %s: PASSED (groups: %v)\n", tt.name, nGot)
			}
		})
	}
}

func BenchmarkGroupAnagrams(b *testing.B) {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	for i := 0; i < b.N; i++ {
		groupAnagrams(input)
	}
}

func BenchmarkGroupAnagramsLarge(b *testing.B) {
	// Deterministic large input with many repeated anagram classes
	const n = 20000
	input := make([]string, 0, n)
	patterns := []string{"aabbcc", "abcabc", "baccab", "zzxyyx", "night", "thing", "listen", "silent"}
	for i := 0; i < n; i++ {
		p := patterns[i%len(patterns)]
		// Slightly vary with index to create multiple classes but still many anagrams
		switch i % 4 {
		case 0:
			input = append(input, p)
		case 1:
			input = append(input, string([]byte{p[1], p[0]})+p[2:])
		case 2:
			input = append(input, p[len(p)/2:]+p[:len(p)/2])
		default:
			input = append(input, p)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		groupAnagrams(input)
	}
}
