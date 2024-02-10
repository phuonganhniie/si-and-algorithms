package leetcode_647

import "fmt"

func CountSubstrings(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	count := 0

	// Single character palindromes
	for i := 0; i < n; i++ {
		dp[i][i] = true
		count++
	}
	fmt.Println("After counting single-character palindromes:", count)

	// Debug print of the DP table after single and two-character checks
	// fmt.Println("DP table after single and two-character palindrome checks:")
	// printDPTable(dp)

	// Two character palindromes
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			fmt.Println("s[i]: ", s[i])
			fmt.Println("s[i+1]: ", s[i+1])
			dp[i][i+1] = true
			count++
		}
	}
	fmt.Println("After counting two-character palindromes:", count)

	// Debug print of the DP table after single and two-character checks
	// fmt.Println("DP table after single and two-character palindrome checks:")
	// printDPTable(dp)

	// Palindromes longer than 2 characters
	for length := 3; length <= n; length++ {
		for i := 0; i < n-length+1; i++ {
			j := i + length - 1
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				count++
			}
		}
	}

	// Final debug print of the complete DP table
	// fmt.Println("Final DP table:")
	// printDPTable(dp)

	return count
}

// printDPTable prints the DP table for debugging purposes.
func printDPTable(dp [][]bool) {
	for _, row := range dp {
		for _, val := range row {
			if val {
				fmt.Print("T ")
			} else {
				fmt.Print("F ")
			}
		}
		fmt.Println()
	}
	fmt.Println("---")
}
