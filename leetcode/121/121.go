package leetcode_121

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	minPrice := prices[0]
	maximumProfit := 0
	for _, price := range prices {
		minPrice = min(price, minPrice)
		profit := price - minPrice
		maximumProfit = max(profit, maximumProfit)
	}
	return maximumProfit
}
