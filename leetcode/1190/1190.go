package leetcode_1190

/*
[Medium] 1190. Reverse Substrings Between Each Pair of Parentheses
https://leetcode.com/problems/reverse-substrings-between-each-pair-of-parentheses/description/
Created: 2024-07-11
Done   : 20 mins 46s
Attempt: 1
---------------------NOTE---------------------
Time: O(n^2)
Space: O(n)
Approach: Stack + 2 pointers
- Cái core là phải xác định được chỗ bắt đầu và kết thúc của ngoặc để mà reverse chuỗi trong ngoặc
*/
func reverseParentheses(s string) string {
	stack := []int{} // parentheses indices
	rs := []byte{}

	for _, char := range s {
		if char == '(' {
			stack = append(stack, len(rs))
		} else if char == ')' {
			start := stack[len(stack)-1] // start parentheses index
			stack = stack[:len(stack)-1]
			reverseString(rs[start:])
		} else {
			rs = append(rs, byte(char))
		}
	}
	return string(rs)
}

func reverseString(arr []byte) {
	start, end := 0, len(arr)-1
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}
