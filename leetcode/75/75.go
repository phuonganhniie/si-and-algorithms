package leetcode_75

/*
[Medium] 75. Sort Colors
https://leetcode.com/problems/sort-colors/description/
Created: 2024-06-12
Done   : 10 mins 41 seconds
Attempt: 3
---------------------NOTE---------------------
Time    : O(n log n)
Space   : O(1)
Approach: Heap Sort
*/
func sortColors(nums []int) []int {
	n := len(nums)

	// xây dựng max heap
	// bắt đầu từ node con của root
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, n, i)
	}

	for i := n - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]

		// tiếp tục heapify từ phần tử đầu tiên với mảng đã giảm đi phần tử cuối
		// với mục đích xây dựng lại max-heap để đảm bảo phần tử root luôn lớn nhất
		heapify(nums, i, 0)
	}

	result := nums
	return result
}

func heapify(nums []int, n, i int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i

	if left < n && nums[left] > nums[largest] {
		largest = left
	}

	if right < n && nums[right] > nums[largest] {
		largest = right
	}

	// kiểm tra xem lúc này node thứ largest có phải node gốc không
	// nếu không, đổi chỗ cho nhau và tiếp tục heapify
	// (nghĩa là node con bên trái hoặc node con bên phải đang có giá trị lớn hơn node gốc)
	if largest != i {
		nums[largest], nums[i] = nums[i], nums[largest]
		heapify(nums, n, largest)
	}
}

/*
Done   : 8 mins 54 seconds
Attempt: 2
---------------------NOTE---------------------
Time    : O(n)
Space   : O(1)
Approach: Two (actual three) Pointers (Dutch National Flag algo)
*/
func sortColors2(nums []int) []int {
	low, mid, high := 0, 0, len(nums)-1

	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}
	result := nums
	return result
}
