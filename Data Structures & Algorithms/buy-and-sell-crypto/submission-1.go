func maxProfit(prices []int) int {
	maxProfit := 0

	for i:=1; i<len(prices); i++ {
		if prices[i-1] < prices[i] {
			if prices[i]-prices[i-1] > maxProfit {
				maxProfit = prices[i]-prices[i-1]
			}
			prices[i] = prices[i-1]
		} 
	}

	return maxProfit
}
