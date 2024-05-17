package leetcode_412

import "fmt"

func helper(i, k int) bool {
	return i%k == 0
}

func FizzBuzz(n int) []string {
	rs := []string{}
	for i := 1; i <= n; i++ {
		if helper(i, 3) && helper(i, 5) {
			rs = append(rs, "FizzBuzz")
			continue
		}
		if helper(i, 3) {
			rs = append(rs, "Fizz")
			continue
		}
		if helper(i, 5) {
			rs = append(rs, "Buzz")
			continue
		}
		rs = append(rs, fmt.Sprintf("%v", i))
	}
	return rs
}
