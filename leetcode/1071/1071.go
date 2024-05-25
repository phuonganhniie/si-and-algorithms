package leetcode_1071

func gcdOfStrings(str1 string, str2 string) string {
	lengthGCD := lenPotentialGCD(len(str1), len(str2))
	potentialStr := str1[:lengthGCD]
	if isDivisible(str1, potentialStr) && isDivisible(str2, potentialStr) {
		return potentialStr
	}
	return ""
}

func lenPotentialGCD(lenStr1, lenStr2 int) int {
	for lenStr2 != 0 {
		lenStr1, lenStr2 = lenStr2, lenStr1%lenStr2
	}
	return lenStr1
}

func isDivisible(str, potentialStr string) bool {
	if len(str)%len(potentialStr) != 0 {
		return false
	}

	for i := 0; i < len(str); i += len(potentialStr) {
		allCharsStr := str[i : i+len(potentialStr)]
		if allCharsStr != potentialStr {
			return false
		}
	}
	return true
}
