package sorting

/**
 * Time: O(n + k)
 * Space: O(n)
 */
func partition012(arr []int) []int {
	maxNum := 2
	countArr := make([]int, maxNum+1, maxNum+1)

	for _, num := range arr {
		countArr[num]++
	}

	rs := []int{}
	for i := 0; i <= maxNum; i++ {
		for range countArr[i] {
			rs = append(rs, i)
		}
	}
	return rs
}

/**
 * Time: O(n)
 * Space: O(1)
 */
func partition012Optimize(arr []int) []int {
	start, mid, end := 0, 0, len(arr)-1

	for mid <= end {
		switch arr[mid] {
		case 0:
			arr[start], arr[mid] = arr[mid], arr[start]
			start++
			mid++
		case 1:
			mid++
		case 2:
			arr[mid], arr[end] = arr[end], arr[mid]
			end--
		}
	}
	return arr
}
