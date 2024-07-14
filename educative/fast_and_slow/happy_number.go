package fastandslow

func sumOfSquareDigits(number int) int {
	totalSum := 0

	for number > 0 {
		digit := number % 10
		number = number / 10
		totalSum += digit * digit
	}
	return totalSum
}

func isHappy(num int) bool {
	slowPointer := num
	fastPointer := sumOfSquareDigits(num)

	for fastPointer != 1 && slowPointer != fastPointer {
		slowPointer = sumOfSquareDigits(slowPointer)
		fastPointer = sumOfSquareDigits(sumOfSquareDigits(fastPointer))
	}

	if fastPointer == 1 {
		return true
	}
	return false
}
