package leetcode_2540

/**
 * Two Pointer approach
 * Complexity: O(n + m)
 * Where: n and m are the lengths of the two arrays.
 * This complexity arises because each pointer can move at most n or m times, respectively, making the total number of steps linearly dependent on the sizes of both arrays.
 */
// func GetCommon(nums1 []int, nums2 []int) int {
// 	i, j := 0, 0

// 	for i < len(nums1) && j < len(nums2) {
// 		if nums1[i] == nums2[j] {
// 			return nums1[i]
// 		}

// 		if nums1[i] < nums2[j] {
// 			i++
// 		} else {
// 			j++
// 		}
// 	}

// 	return -1
// }

/**
 * Binary Search approach
 *
 */
func binarySearch(arr []int, target int) bool {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return true
		}

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

func GetCommon(nums1 []int, nums2 []int) int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	for _, num := range nums1 {
		if binarySearch(nums2, num) {
			return num
		}
	}
	return -1
}
