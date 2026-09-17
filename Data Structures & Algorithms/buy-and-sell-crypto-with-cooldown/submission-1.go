func maxProfit(prices []int) int {
	memo := make(map[int]int)
    return recurse(prices, 0, memo)
}

func recurse(prices []int, index int, memo map[int]int) int {
	maxValue := 0

	for i:= index; i<len(prices); i++ {
		for j:=i+1; j < len(prices); j++ {
			if prices[j]-prices[i] > 0 {
				currValue := prices[j]-prices[i]
				if j+2 < len(prices) {
					if val, ok := memo[j+2]; ok {
						currValue += val
					} else {
						val1 := recurse(prices, j+2, memo)
						currValue += val1
						memo[j+2] = val1
					}
				}
				if currValue > maxValue {
					maxValue = currValue
				}
			}
		}
	}

	return maxValue
}
