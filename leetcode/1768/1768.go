package leetcode_1768

func MergeAlternately(word1 string, word2 string) string {
	i, j := 0, 0
	var mergedStr []byte
	for i < len(word1) && j < len(word2) {
		mergedStr = append(mergedStr, word1[i])
		mergedStr = append(mergedStr, word2[j])
		i++
		j++
	}

	if i < len(word1) {
		mergedStr = append(mergedStr, word1[i:]...)
	}
	if j < len(word2) {
		mergedStr = append(mergedStr, word2[j:]...)
	}

	return string(mergedStr)
}
