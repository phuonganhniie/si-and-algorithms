package stack

import (
	"strconv"
	"unicode"
)

// TODO: Retry!!!
func calculator(expression string) int {
	number, signValue, result := 0, 1, 0
	calculatorStack := make([]int, len(expression))
	signMap := map[rune]int{
		'+': 1,
		'-': -1,
	}

	for _, char := range expression {
		if unicode.IsDigit(char) {
			temp, _ := strconv.Atoi(string(char))
			number = number*10 + temp
		}

		if sign, found := signMap[char]; found {
			result += number * signValue
			signValue = sign
			number = 0
		}

		if char == '(' {
			calculatorStack = append(calculatorStack, result, signValue)
			result, signValue = 0, 1
		}

		if char == ')' {
			result += signValue * number
			number = 0

			// pop from stack
			if len(calculatorStack) > 0 {
				signValue = calculatorStack[len(calculatorStack)-1]
				calculatorStack = calculatorStack[:len(calculatorStack)-1]
			}

			if len(calculatorStack) > 0 {
				prevResult := calculatorStack[len(calculatorStack)-1]
				calculatorStack = calculatorStack[:len(calculatorStack)-1]

				result = prevResult + result*signValue
			}
		}
	}
	return result + number*signValue
}
