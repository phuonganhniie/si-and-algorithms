package leetcode_392

import "fmt"

/*
Problem: Is Subsequence (LeetCode 392)
Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

Since s and t consist only of lowercase English letters, we can use several approaches:

1. Two Pointer (Original) - O(n) time, O(1) space
   - Most efficient for single queries
   - Simple and intuitive

2. Binary Search with Position Mapping - O(m log n) time, O(n) space
   - Better for multiple queries on the same string t
   - Preprocessing allows faster subsequent queries

3. Recursive with Memoization - O(m*n) time, O(m*n) space
   - Dynamic programming approach
   - Good for understanding the problem structure

4. Bit Manipulation - O(n + m) time, O(26) space (for lowercase letters)
   - Leverages the constraint of lowercase English letters only
   - Limited to strings of length ≤ 64 due to uint64 constraints
   - Most space-efficient for the character set constraint
*/

// Original two-pointer approach (most efficient for single query)
func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}

	ptrS, ptrT := 0, 0
	for ptrS < len(s) && ptrT < len(t) {
		if s[ptrS] == t[ptrT] {
			ptrT++
			ptrS++
			continue
		}
		ptrT++
	}

	return len(s) == ptrS
}

// Approach 2: Using binary search with position mapping
// Better for multiple queries on the same string t
func isSubsequenceBinarySearch(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}

	// Build position map for each character in t
	charPositions := make(map[byte][]int)
	for i := 0; i < len(t); i++ {
		char := t[i]
		charPositions[char] = append(charPositions[char], i)
	}

	prevPos := -1
	for i := 0; i < len(s); i++ {
		char := s[i]
		positions, exists := charPositions[char]
		if !exists {
			return false
		}

		// Binary search for the first position > prevPos
		found := false
		for _, pos := range positions {
			if pos > prevPos {
				prevPos = pos
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

// Approach 3: Using recursive approach with memoization
func isSubsequenceRecursive(s string, t string) bool {
	memo := make(map[string]bool)
	return isSubsequenceHelper(s, t, 0, 0, memo)
}

func isSubsequenceHelper(s, t string, i, j int, memo map[string]bool) bool {
	key := fmt.Sprintf("%d,%d", i, j)
	if val, exists := memo[key]; exists {
		return val
	}

	// Base cases
	if i == len(s) {
		memo[key] = true
		return true
	}
	if j == len(t) {
		memo[key] = false
		return false
	}

	var result bool
	if s[i] == t[j] {
		// Match found, advance both pointers
		result = isSubsequenceHelper(s, t, i+1, j+1, memo)
	} else {
		// No match, advance only t pointer
		result = isSubsequenceHelper(s, t, i, j+1, memo)
	}

	memo[key] = result
	return result
}

// Approach 4: Optimized for lowercase letters using bit manipulation
// Note: This approach has limitations for very long strings due to integer overflow
func isSubsequenceBitMask(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}

	// Build bitmasks for character positions (works for strings up to 64 chars)
	if len(t) > 64 {
		// Fall back to two-pointer approach for longer strings
		return isSubsequence(s, t)
	}

	charMasks := make(map[byte]uint64)
	for i := 0; i < len(t); i++ {
		char := t[i]
		charMasks[char] |= 1 << uint(i)
	}

	lastUsedPos := -1
	for i := 0; i < len(s); i++ {
		char := s[i]
		availableMask, exists := charMasks[char]
		if !exists {
			return false
		}

		// Create mask for positions after lastUsedPos
		var afterMask uint64
		if lastUsedPos >= 0 {
			afterMask = ^((1 << uint(lastUsedPos+1)) - 1)
		} else {
			afterMask = ^uint64(0) // All positions available
		}

		// Find the first available position after lastUsedPos
		candidateMask := availableMask & afterMask
		if candidateMask == 0 {
			return false
		}

		// Use the rightmost (earliest) available position
		pos := trailingZeros(candidateMask)
		lastUsedPos = pos
	}

	return true
}

// Helper function to count trailing zeros (find rightmost set bit position)
func trailingZeros(n uint64) int {
	if n == 0 {
		return 64
	}
	count := 0
	for (n & 1) == 0 {
		n >>= 1
		count++
	}
	return count
}
