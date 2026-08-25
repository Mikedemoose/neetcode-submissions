func maxProfit(prices []int) int {

	minPriceHistory := make([]int, len(prices))

	minPriceHistory[0] = prices[0]
	maxProfit := 0

	for i:=1; i<len(prices); i++ {
		if minPriceHistory[i-1] < prices[i] {
			minPriceHistory[i] = minPriceHistory[i-1]
			if prices[i]-minPriceHistory[i] > maxProfit {
				maxProfit = prices[i]-minPriceHistory[i]
			}
		} else {
			minPriceHistory[i] = prices[i]
		}
	}

	return maxProfit
}
