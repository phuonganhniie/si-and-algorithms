package slidingwindow

import (
	"fmt"

	"github.com/phuonganhniie/educative/helper"
)

func findMaxSlidingWindow(nums []int, windowSize int) []int {
	if len(nums) == 0 || windowSize == 0 {
		return []int{}
	}

	deque := &helper.Deque{}
	result := []int{}

	for i := 0; i < len(nums); i++ {
		fmt.Printf("*** The %d iteration\n", i+1)
		windowStart := i - windowSize + 1

		// Print the current window
		if windowSize-1 <= i {
			fmt.Printf("Window [%d->%d]: %v\n", windowStart, i, nums[windowStart:i+1])
		}

		// Remove elements smaller than nums[i]
		for len(deque.Elements) > 0 && nums[deque.Back()] < nums[i] {
			deque.Pop()
		}

		// Remove the elements out of current window
		if len(deque.Elements) > 0 && deque.Front() < windowStart {
			deque.PopLeft()
		}

		// Add current element at the end of the queue
		deque.Append(i)

		// First element in the deque is the maximum of the current window
		if windowSize-1 <= i {
			fmt.Printf("Current indices in deque: %v\n", deque.Elements)
			result = append(result, nums[deque.Front()])
			fmt.Printf("Current result: %v\n", result)
		}
	}
	return result
}
