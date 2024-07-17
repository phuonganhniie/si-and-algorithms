package leetcode_187

/*
[Medium] 187. Repeated DNA Sequences
https://leetcode.com/problems/repeated-dna-sequences/description/
Created: 2024-07-18
Done   : 11 mins 47s
Attempt: 1
---------------------NOTE---------------------
Time: O(N)
Space: O(N)
Approach: Counting repeated
*/
func findRepeatedDnaSequences(s string) []string {
	const k = 10

	if len(s) < k {
		return []string{}
	}

	result := []string{}
	countRepeated := make(map[string]int)

	for i := 0; i <= len(s)-k; i++ {
		subStr := s[i : i+k]

		countRepeated[subStr]++

		if countRepeated[subStr] == 2 {
			result = append(result, subStr)
		}
	}
	return result
}
