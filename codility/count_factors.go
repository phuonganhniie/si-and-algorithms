package codility

func SolutionLesson10_1(N int) int {
	count := 0
	i := 1

	for i*i <= N {
		if N%i == 0 {
			count += 2
			if i*i == N {
				count -= 1
			}
		}
		i += 1
	}
	return count
}
