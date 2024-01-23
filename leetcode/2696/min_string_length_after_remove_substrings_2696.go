package leetcode

func MinLength2(s string) int {
	stack := []rune{}

	for i := range s {
		if len(stack) > 0 && ((stack[len(stack)-1] == 'A' && s[i] == 'B') || (stack[len(stack)-1] == 'C' && s[i] == 'D')) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, int32(s[i]))
		}
	}
	return len(stack)
}

func MinLength(s string) int {
	stack := []byte{}

	for _, ch := range s {
		if len(stack) > 0 {
			topElement := stack[len(stack)-1]

			if topElement == 'A' && ch == 'B' || topElement == 'C' && ch == 'D' {
				stack = stack[:len(stack)-1]
				continue
			}
		}
		stack = append(stack, byte(ch))
	}

	return len(stack)
}
