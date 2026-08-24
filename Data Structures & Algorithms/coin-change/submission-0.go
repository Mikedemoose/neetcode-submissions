func coinChange(coins []int, amount int) int {
    memo := make([]int, amount+1)
	memo[0] = 0

	for i := 1; i <= amount; i++ {
		memo[i] = -1
		for _, coin := range coins {
			if coin > i {
				continue
			}

			if memo[i-coin] != -1 {
				if memo[i] == -1 {
					memo[i] = memo[i-coin] + 1
				} else if memo[i-coin] + 1 < memo[i] {
					memo[i] = memo[i-coin] + 1
				}
			}
		}
	}

	return memo[amount]
}
