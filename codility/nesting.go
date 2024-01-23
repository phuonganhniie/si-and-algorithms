package codility

func SolutionLesson7_3(S string) int {
	stack := []byte{}

	for _, char := range S {
		if char == '(' {
			stack = append(stack, byte(char))
		}

		if char == ')' {
			if len(stack) == 0 {
				return 0
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		return 1
	}
	return 0
}
