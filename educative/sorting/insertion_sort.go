package sorting

// Visualization: https://www.youtube.com/watch?v=8mJ-OhcfpYg&t=373s
func InsertionSort(arr []int) []int {
	n := len(arr)
	var temp int

	for i := 1; i < n; i++ {
		temp = arr[i]
		for j := i - 1; j >= 0 && arr[j] > temp; j-- {
			arr[j+1] = arr[j]
			arr[j] = temp
		}
	}
	return arr
}
