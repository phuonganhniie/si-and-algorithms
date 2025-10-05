package leetcode_347

import "sort"

func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int) //key: number - value: count

	for _, num := range nums {
		countMap[num]++
	}

	bucket := make([][]int, len(nums)+1)
	for num, count := range countMap {
		bucket[count] = append(bucket[count], num)
	}

	result := make([]int, 0, k)
	for i := len(bucket) - 1; i >= 0 && len(result) < k; i-- {
		if len(bucket[i]) == 0 {
			continue
		}

		for _, num := range bucket[i] {
			result = append(result, num)
			if len(result) == k {
				break
			}
		}
	}
	return result
}

func topKFrequentSlower(nums []int, k int) []int {
	countMap := make(map[int]int) //key: number - value: count

	for _, num := range nums {
		countMap[num]++
	}

	type pair struct {
		number int
		count  int
	}
	pairs := []pair{}
	for num, count := range countMap {
		pairs = append(pairs, pair{
			number: num,
			count:  count,
		})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})

	result := make([]int, 0, k)
	for _, pair := range pairs {
		if len(result) >= k {
			break
		}
		result = append(result, pair.number)
	}
	return result
}
