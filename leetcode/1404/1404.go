package leetcode_1404

import (
	"math/big"
)

func numSteps(s string) int {
	steps := 0
	zero, one, two := big.NewInt(0), big.NewInt(1), big.NewInt(2)
	number := binaryToDecimal(s)

	for number.Cmp(one) > 0 {
		if new(big.Int).Mod(number, two).Cmp(zero) == 0 {
			number.Div(number, two)
		} else {
			number.Add(number, one)
		}
		steps++
	}
	return steps
}

func binaryToDecimal(binaryString string) *big.Int {
	decimalNumber := new(big.Int)

	for i := 0; i < len(binaryString); i++ {
		decimalNumber.Lsh(decimalNumber, 1)
		if binaryString[i] == '1' {
			decimalNumber.Or(decimalNumber, big.NewInt(1))
		}
	}
	return decimalNumber
}

// Chat GPT solution
func numStepsGPT(s string) int {
	steps := 0
	for len(s) > 1 {
		if s[len(s)-1] == '0' {
			s = divideByTwo(s)
		} else {
			s = addOne(s)
		}
		steps++
	}
	return steps
}

func divideByTwo(s string) string {
	return s[:len(s)-1]
}

func addOne(s string) string {
	carry := true
	result := []byte(s)
	for i := len(s) - 1; i >= 0; i-- {
		if carry {
			if s[i] == '0' {
				result[i] = '1'
				carry = false
			} else {
				result[i] = '0'
			}
		}
	}
	if carry {
		result = append([]byte{'1'}, result...)
	}
	return string(result)
}
