package leetcode_1002

func commonChars(words []string) []string {
	commonCharsMap := make(map[rune]int)
	for _, char := range words[0] {
		commonCharsMap[char]++
	}

	for _, s := range words[1:] {
		tempMap := make(map[rune]int)
		for _, char := range s {
			tempMap[char]++
		}

		for char := range commonCharsMap {
			if count, exist := tempMap[char]; exist {
				commonCharsMap[char] = min(count, commonCharsMap[char])
			} else {
				delete(commonCharsMap, char)
			}
		}
	}

	result := make([]string, 0, len(words[0]))
	for char, count := range commonCharsMap {
		for i := 0; i < count; i++ {
			result = append(result, string(char))
		}
	}
	return result
}

func commonCharsGemini(words []string) []string {
	// Initialize counts for lowercase letters (a-z)
	minCounts := make([]int, 26)
	for i := range minCounts {
		minCounts[i] = 101 // Initialize with a value greater than max word length
	}

	// Iterate through words and update min counts
	for _, word := range words {
		charCounts := make([]int, 26) // Count for current word
		for _, char := range word {
			charCounts[char-'a']++ // Increment count for this character
		}
		for i := range minCounts {
			if charCounts[i] < minCounts[i] {
				minCounts[i] = charCounts[i] // Update minimum count
			}
		}
	}

	// Construct result based on min counts
	result := make([]string, 0, len(words[0]))
	for i, count := range minCounts {
		for j := 0; j < count; j++ {
			result = append(result, string(rune('a'+i))) // Convert index back to character
		}
	}
	return result
}
