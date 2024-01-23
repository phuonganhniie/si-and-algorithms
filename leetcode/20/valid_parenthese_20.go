package leetcode

import "fmt"

func IsValid(s string) bool {
	if len(s) == 0 {
		return false
	}

	validChars := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}

	stack := []rune{}
	for _, r := range s {
		fmt.Println("rune is: ", r)
		if _, ok := validChars[r]; ok {
			stack = append(stack, r)
		} else if len(stack) == 0 || validChars[stack[len(stack)-1]] != r {
			fmt.Println("stack in else if: ", validChars[stack[len(stack)-1]])
			return false
		} else {
			stack = stack[:len(stack)-1]
			fmt.Println("stack in else: ", stack)
		}
	}
	return len(stack) == 0
}

func ShiftArrayRight(arr []int) {
	if arr == nil || len(arr) <= 1 {
		return
	}

	lastElement := arr[len(arr)-1]
	for i := len(arr) - 2; i >= 0; i-- {
		arr[i+1] = arr[i]
	}
	arr[0] = lastElement
}
