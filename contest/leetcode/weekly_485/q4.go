package weekly485

func smallestString(s string) string {
	// Get unique characters and track what's remaining
	remaining := make(map[rune]int)
	for _, ch := range s {
		remaining[ch]++
	}

	stack := []rune{}
	inStack := make(map[rune]int)
	used := make(map[rune]bool) // Track if we've added this char type at least once

	for _, ch := range s {
		// Pop larger characters if they appear later
		for len(stack) > 0 && stack[len(stack)-1] > ch {
			top := stack[len(stack)-1]
			if remaining[top] > 0 || inStack[top] > 1 {
				stack = stack[:len(stack)-1]
				inStack[top]--
				if inStack[top] == 0 {
					used[top] = false
				}
			} else {
				break
			}
		}

		// Add current character if:
		// 1. Not used yet (first occurrence), OR
		// 2. Already used AND there exists a larger char we haven't used yet
		shouldAdd := !used[ch]
		if !shouldAdd {
			// Check if any larger character hasn't been used yet
			for larger := ch + 1; larger <= 'z'; larger++ {
				if remaining[larger] > 0 && !used[larger] {
					shouldAdd = true
					break
				}
			}
		}

		if shouldAdd {
			stack = append(stack, ch)
			inStack[ch]++
			used[ch] = true
		}

		remaining[ch]--
	}

	return string(stack)
}
