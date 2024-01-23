package codility

import (
	"fmt"
	"sort"
)

/**
Problem: Given an array of integers. Find maximum number of non-intersecting elements of length 2(adjacent elements) having same sum.

Eg: arr[] = {10,1,3,1,2,2,1,0,4}
ouput: 3
Explanation: non-intersecting adjacent pairs having sum = 4 are (1,3),(2,2) and (0,4)

Eg: arr[] = { 5,3,1,3,2,3}
Output: 1
Explanation: no adjacent pairs have same sum

Eg: arr[] ={9,9,9,9}
Output: 2
*/

func MaxNonIntersectingPairs(arr []int) int {
	fmt.Printf("Input array: %v\n", arr)
	if len(arr) < 2 {
		return 0
	}

	sumMap := make(map[int][]int)
	for i := 0; i < len(arr)-1; i++ {
		sum := arr[i] + arr[i+1]
		sumMap[sum] = append(sumMap[sum], i)
	}

	fmt.Println("Sum Map:", sumMap)

	maxPairs := 0
	for sum, indices := range sumMap {
		fmt.Printf("Processing sum = %d with indices %v\n", sum, indices)

		count := 0
		prevIndex := -1
		sort.Ints(indices)
		fmt.Printf("Sorted indices for sum %d: %v\n", sum, indices)

		for _, index := range indices {
			if index > prevIndex {
				count++
				prevIndex = index + 1
				fmt.Printf("Selected pair at index %d, count = %d\n", index, count)
			} else {
				fmt.Printf("Skipping pair at index %d\n", index)
			}
		}

		if count > maxPairs {
			maxPairs = count
		}
	}

	return maxPairs
}

func RewriteMaxNonIntersectingPairs(arr []int) int {
	if len(arr) < 2 {
		return 0
	}

	maxPairs := 0
	sumMap := make(map[int][]int)
	for i := 0; i < len(arr)-1; i++ {
		sum := arr[i] + arr[i+1]
		sumMap[sum] = append(sumMap[sum], i)
	}

	for _, indices := range sumMap {
		sort.Ints(indices)

		count := 0
		prevIndex := -1
		for _, index := range indices {
			if index > prevIndex {
				count++
				prevIndex = index + 1
			}
		}
		if count > maxPairs {
			maxPairs = count
		}
	}
	return maxPairs
}
