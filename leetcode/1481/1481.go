/**
 * 1481. Least Number of Unique Integers after K Removals
 * https://leetcode.com/problems/least-number-of-unique-integers-after-k-removals/description/?envType=daily-question&envId=2024-02-16
 */

package leetcode_1481

import "sort"

func FindLeastNumOfUniqueInts(arr []int, k int) int {
	n := len(arr)
	freqMap := make(map[int]int, n)

	for _, num := range arr {
		if _, ok := freqMap[num]; ok {
			freqMap[num] += 1
		} else {
			freqMap[num] = 1
		}
	}

	frequencies := make([]int, 0, len(freqMap))
	for _, freq := range freqMap {
		frequencies = append(frequencies, freq)
	}

	sort.Ints(frequencies)

	distinctInts := len(freqMap)
	for _, freq := range frequencies {
		if k >= freq {
			k = k - freq
			distinctInts--
		} else {
			break
		}
	}

	return distinctInts
}
