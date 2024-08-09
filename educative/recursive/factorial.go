package recursive

func factorial(i int) int {
	if i <= 1 {
		return i
	}

	rs := factorial(i - 1)
	rs *= i
	return rs
}
