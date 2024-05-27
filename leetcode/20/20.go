package leetcode_20

func isValid(s string) bool {
	if len(s) <= 1 {
		return false
	}

	validParenthesesMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	stack := []rune{}
	for _, char := range s {
		if _, exist := validParenthesesMap[char]; exist {
			stack = append(stack, char)
			continue
		}

		if len(stack) == 0 && validParenthesesMap[char] != char {
			return false
		}

		if len(stack) > 0 {
			lastChar := stack[len(stack)-1]
			if validParenthesesMap[lastChar] != char {
				break
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
