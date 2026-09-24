func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	costMap := make(map[int][]QueueItem)
	for _, flight := range flights {
		costMap[flight[0]] = append(costMap[flight[0]], QueueItem{
			curr: flight[1],
			cost: flight[2],
		})
	}

	queue := []QueueItem { QueueItem { curr: src }}

	cheapestCost := make([]int, 100)
	for i := range 100 {
		cheapestCost[i] = -1
	}

	minCost := -1

	for range k+2 {
		newQueue := make([]QueueItem, 0)
		for _, item := range queue {
			if cheapestCost[item.curr] != -1 && cheapestCost[item.curr] < item.cost {
				continue
			} else {
				cheapestCost[item.curr] = item.cost
			}


			if item.curr == dst {
				if minCost == -1 || item.cost < minCost {
					minCost = item.cost
				}
			}
			next := costMap[item.curr]

			for _, nextStop := range next {
				newQueue = append(newQueue, QueueItem{
					curr: nextStop.curr,
					cost: item.cost + nextStop.cost,
				})

			}

		}
		queue = newQueue
	}

	return minCost
}


type QueueItem struct {
	curr int
	cost int
}
