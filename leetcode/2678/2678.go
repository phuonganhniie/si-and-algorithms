package leetcode_2678

/*
[Medium] 2678. Number of Senior Citizens
https://leetcode.com/problems/number-of-senior-citizens/description/?envType=daily-question&envId=2024-08-01
Created: 01-08-2024
Done   : 8m50s
Attempt: 1
---------------------NOTE---------------------
Time: O(N)
Space: O(1)
Intuition: Slipt string or convert character and multiply to 10
*/
func countSeniors(details []string) int {
	count := 0

	for _, info := range details {
		ageTen := info[11] - '0'
		ageOne := info[12] - '0'

		age := 0
		age = int(ageTen)*10 + int(ageOne)

		if age > 60 {
			count++
		}
	}

	return count
}
