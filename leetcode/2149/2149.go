/**
 * 2149. Rearrange Array Elements by Sign
 * https://leetcode.com/problems/rearrange-array-elements-by-sign/description/?envType=daily-question&envId=2024-02-14
 */

package leetcode_2149

func RearrangeArray(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	posPointer, negPointer := 0, 1 // Start posIndex at 0 for positive numbers and negIndex at 1 for negative numbers

	for _, num := range nums {
		if num > 0 {
			// Place positive number at posPointer and move posPointer forward by 2
			result[posPointer] = num
			posPointer += 2
		}

		if num < 0 {
			// Place negative number at negPointer and move negPointer forward by 2
			result[negPointer] = num
			negPointer += 2
		}
	}

	return result
}
