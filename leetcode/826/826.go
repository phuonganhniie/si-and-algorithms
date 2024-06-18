package leetcode_826

import "sort"

/*
[Medium] 862. Most Profit Assigning Work
https://leetcode.com/problems/most-profit-assigning-work/description/
Created: 2024-06-18
Done   : 50 mins 09 seconds
Attempt: 4
---------------------NOTE---------------------
Time: O(nlogn)
Space: O(n)
Approach: Two Pointers + Sorting
Intuition:
- Sort worker và pairing độ khó và tiền công của một job, sau đó sort paired jobs.
- Loop qua từng ability của worker, kiểm tra xem nếu ability lớn hơn độ khó của job:
+ Tiếp tục kiểm tra profit của job có lớn hơn max profit không.
+ Nếu có thì cập nhật lại, còn không thì tiếp tục xét job khó hơn.
- Nếu ability bé hơn độ khó của job: cộng tổng profit với max profit
- Return total profit
*/
type Job struct {
	difficulty, profit int
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	sort.Ints(worker)

	jobs := make([]Job, len(difficulty))
	for i := range difficulty {
		jobs[i] = Job{difficulty: difficulty[i], profit: profit[i]}
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].difficulty < jobs[j].difficulty
	})

	totalProfit, maxProfit, jobPtr := 0, 0, 0
	for _, ability := range worker {
		for jobPtr < len(jobs) && ability >= jobs[jobPtr].difficulty {
			currentProfit := jobs[jobPtr].profit
			if currentProfit > maxProfit {
				maxProfit = currentProfit
			}
			jobPtr++
		}
		totalProfit += maxProfit
	}
	return totalProfit
}
