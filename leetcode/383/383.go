package leetcode_383

func CanConstruct(ransomNote string, magazine string) bool {
	constructMap := make(map[rune]int)

	for _, ch := range magazine {
		constructMap[ch]++
	}

	for _, ch := range ransomNote {
		if qty, ok := constructMap[ch]; !ok || qty == 0 {
			return false
		}
		constructMap[ch]--
	}

	return true
}
