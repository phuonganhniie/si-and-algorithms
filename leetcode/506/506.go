package leetcode_506

import (
	"fmt"
	"sort"
)

func findRelativeRanks(score []int) []string {
	lenScore := len(score)
	ans := make([]string, lenScore)
	scoreWithInx := make([][2]int, lenScore)

	for i, s := range score {
		scoreWithInx[i] = [2]int{s, i}
	}

	sort.Slice(scoreWithInx, func(i, j int) bool {
		return scoreWithInx[i][0] > scoreWithInx[j][0]
	})

	rank := []string{"Gold Medal", "Silver Medal", "Bronze Medal"}
	for i := 0; i < lenScore; i++ {
		index := scoreWithInx[i][1]
		if i < 3 {
			ans[index] = rank[i]
		} else {
			ans[index] = fmt.Sprintf("%d", i+1)
		}
	}

	return ans
}

func findRelativeRanksGPT(score []int) []string {
	// Step 1: Create an array of tuples (score, index)
	n := len(score)
	sortedScores := make([][2]int, n)
	for i := range score {
		sortedScores[i] = [2]int{score[i], i}
	}

	// Step 2: Sort the scores in descending order
	sort.Slice(sortedScores, func(i, j int) bool {
		return sortedScores[i][0] > sortedScores[j][0]
	})

	// Step 3: Create the ranks result array
	ranks := make([]string, n)
	for i, pair := range sortedScores {
		switch i {
		case 0:
			ranks[pair[1]] = "Gold Medal"
		case 1:
			ranks[pair[1]] = "Silver Medal"
		case 2:
			ranks[pair[1]] = "Bronze Medal"
		default:
			ranks[pair[1]] = fmt.Sprintf("%d", i+1)
		}
	}

	return ranks
}
