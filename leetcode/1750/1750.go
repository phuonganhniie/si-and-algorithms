package leetcode_1750

// func MinimumLength(s string) int {
// 	n := len(s)
// 	left := 0
// 	right := n - 1

// 	for left < right {
// 		if s[left] != s[right] {
// 			return right - left + 1
// 		}

// 		for left < right && s[left] == s[left+1] {
// 			left++
// 		}

// 		for left < right && s[right] == s[right-1] {
// 			right--
// 		}

// 		left++
// 		right--
// 	}

// 	if left == right {
// 		return 1
// 	}

// 	return 0
// }

func MinimumLength(s string) int {
	n := len(s)
	left := 0
	right := n - 1

	for left < right {
		leftChar := s[left]
		rightChar := s[right]

		if leftChar == rightChar {
			for left <= right && s[left] == rightChar {
				left++
			}

			for left <= right && s[right] == leftChar {
				right--
			}
		} else {
			break
		}
	}

	return right - left + 1
}
