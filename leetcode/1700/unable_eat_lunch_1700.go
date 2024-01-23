package leetcode

func CountStudents(students []int, sandwiches []int) int {
	unableEatStudents := 0
	for len(students) > 0 {
		if students[0] == sandwiches[0] {
			students = students[1:]
			sandwiches = sandwiches[1:]
			unableEatStudents = 0
		} else {
			unableEatStudents++
		}

		if unableEatStudents >= len(students) {
			break
		}

		students = append(students[1:], students[0])
	}
	return len(students)
}
