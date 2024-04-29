package leetcode_2997

func MinOperations(nums []int, k int) int {
	currentXOR := 0
	for _, num := range nums {
		currentXOR ^= num
	}

	// The required changes in XOR
	targetChange := currentXOR ^ k
	minFlips := 0

	// Count the number of bits that need to be flipped
	for targetChange > 0 {
		if targetChange&1 == 1 {
			minFlips++
		}
		targetChange >>= 1
	}

	return minFlips
}
