package weekly471

func longestBalancedSubstring(s string) int {
	maxLen := 0

	// Try each possible subset of characters
	// We'll use bitmask: bit 0 = include 'a', bit 1 = include 'b', bit 2 = include 'c'
	for mask := 1; mask < 8; mask++ {
		chars := make([]byte, 0, 3)
		if mask&1 != 0 {
			chars = append(chars, 'a')
		}
		if mask&2 != 0 {
			chars = append(chars, 'b')
		}
		if mask&4 != 0 {
			chars = append(chars, 'c')
		}

		length := longestBalancedForChars(s, chars)
		if length > maxLen {
			maxLen = length
		}
	}

	return maxLen
}

// Find longest balanced substring containing exactly the given characters
func longestBalancedForChars(s string, chars []byte) int {
	n := len(s)
	k := len(chars)
	maxLen := 0

	// Map character to index
	charToIdx := make(map[byte]int)
	for i, ch := range chars {
		charToIdx[ch] = i
	}

	// Use state map: state represents difference pattern
	// state = count[0] - count[1], count[1] - count[2], etc.
	// If all chars have same count, all differences are 0
	type State struct {
		d1, d2 int // differences between counts
	}

	// Map from state to earliest position where this state occurred
	statePos := make(map[State]int)

	// Initial state at position -1
	statePos[State{0, 0}] = -1

	count := make([]int, k)
	validStart := 0 // Track where the current valid segment starts

	for i := 0; i < n; i++ {
		// Check if current character is in our set
		idx, exists := charToIdx[s[i]]
		if !exists {
			// Character not in our set, reset everything
			count = make([]int, k)
			statePos = make(map[State]int)
			statePos[State{0, 0}] = i
			validStart = i + 1
			continue
		}

		count[idx]++

		// Calculate current state (difference pattern)
		var currentState State
		if k >= 2 {
			currentState.d1 = count[0] - count[1]
		}
		if k >= 3 {
			currentState.d2 = count[1] - count[2]
		}

		// Check if we've seen this state before
		if prevPos, found := statePos[currentState]; found {
			// From prevPos+1 to i, characters in our set have same count
			// Only consider if the segment starts at or after validStart
			if prevPos >= validStart-1 {
				length := i - prevPos
				if length > maxLen {
					maxLen = length
				}
			}
		} else {
			statePos[currentState] = i
		}
	}

	return maxLen
}
