func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    inf := 1001
	distances := make([]int, n)
	for i := range n {
		distances[i] = inf
	}

	distances[src] = 0

	cheapestPrice := -1

	for range k+1 {
		minDistMap := make(map[int]int)
		for _, flight := range flights {
			dest := flight[1]
			currCost := distances[flight[0]] + flight[2]

			if currCost < distances[dest] {
				if val, ok := minDistMap[dest]; !ok {
					minDistMap[dest] = currCost
				} else {
					minDistMap[dest] = min(currCost, val)
				}
			}
		}
		for k, v := range minDistMap {
			distances[k] = v
		}
	}

	if distances[dst] != inf {
		cheapestPrice = distances[dst]
	}

	return cheapestPrice
}
