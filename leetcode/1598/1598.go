package leetcode_1598

/*
[Medium] 1598. Crawler Log Folder
https://leetcode.com/problems/crawler-log-folder/description/
Created: 2024-07-10
Done   : 10 mins 31s
Attempt: 1
---------------------NOTE---------------------
Time: O(n)
Space: O(1)
Approach: Simulator -> suy luận
*/
func minOperations(logs []string) int {
	level := 0
	for i := 0; i < len(logs); i++ {
		switch logs[i] {
		case "../":
			if level > 0 {
				level--
			}
		case "./":
			continue
		default:
			level++
		}
	}
	return level
}

func minOperations2(logs []string) int {
	level := 0
	actions := map[string]int{
		"../": -1,
		"./":  0,
	}

	for _, log := range logs {
		if action, exist := actions[log]; exist {
			if action == -1 && level > 0 {
				level--
			}
		} else {
			level++
		}
	}
	return level
}
