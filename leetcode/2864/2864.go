package leetcode_2864

func MaximumOddBinaryNumber(s string) string {
	n := len(s)
	firstOne := true
	last, ans := "", ""

	for i := 0; i < n; i++ {
		if firstOne && s[i] == '1' {
			firstOne = !firstOne
			last += string(s[i])
			continue
		}

		if !firstOne && s[i] == '1' {
			ans = string(s[i]) + ans
		} else {
			ans += string(s[i])
		}
	}
	return ans + last
}
