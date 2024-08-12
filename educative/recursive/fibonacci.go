package recursive

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	// rs := fibonacci(n - 1)

	// rs += fibonacci(n - 2)

	// return rs

	return fibonacci(n-1) + fibonacci(n-2)
}
