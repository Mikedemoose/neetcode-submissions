func maxProfit(prices []int) int {
	maxProfit := 0

	for i:=1; i<len(prices); i++ {
		profit := prices[i]-prices[i-1]
		if profit > 0 {
			if profit > maxProfit {
				maxProfit = profit
			}
			prices[i] = prices[i-1]
		} 
	}

	return maxProfit
}
