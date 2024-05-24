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

	for windowEnd := 0; windowEnd < len(nums); windowEnd++ {
		fmt.Printf("*** The %d iteration\n", windowEnd+1)

		windowStart := windowEnd - windowSize + 1

		// Print the current window
		if windowSize-1 <= windowEnd {
			fmt.Printf("Window [%d->%d]: %v\n", windowStart, windowEnd, nums[windowStart:windowEnd+1])
		}

		// Remove elements smaller than nums[i]
		for len(deque.Elements) > 0 && nums[deque.Back()] < nums[windowEnd] {
			deque.Pop()
		}

		// Remove the elements out of current window
		if len(deque.Elements) > 0 && deque.Front() < windowStart {
			deque.PopLeft()
		}

		// Add current element's index at the end of the queue
		deque.Append(windowEnd)

		// First element in the deque is the maximum of the current window
		if windowSize-1 <= windowEnd {
			fmt.Printf("Current indices in deque: %v\n", deque.Elements)
			result = append(result, nums[deque.Front()])
			fmt.Printf("Current result: %v\n", result)
		}
	}
	return result
}
