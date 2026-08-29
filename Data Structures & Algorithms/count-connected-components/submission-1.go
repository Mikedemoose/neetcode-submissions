func countComponents(n int, edges [][]int) int {
	nodeConnectionMap := make(map[int]map[int]struct{})
	for _, edge := range edges {
		if _, ok := nodeConnectionMap[edge[0]]; !ok {
			nodeConnectionMap[edge[0]] = make(map[int]struct{})
		}
		if _, ok := nodeConnectionMap[edge[1]]; !ok {
			nodeConnectionMap[edge[1]] = make(map[int]struct{})
		}
		nodeConnectionMap[edge[0]][edge[1]] = struct{}{}
		nodeConnectionMap[edge[1]][edge[0]] = struct{}{}
	}

	visited := make(map[int]struct{})
	result := 0

	for node := range n {
		if _, ok := visited[node]; ok {
			continue
		}

		result++
		queue := []int{node}


		for len(queue) > 0 {
			newQueue := make([]int, 0)
			for _, item := range queue {
				visited[item] = struct{}{}
				if val, ok := nodeConnectionMap[item]; !ok {
					break
				} else {
					for k, _ := range val {
						if _, ok := visited[k]; !ok {
							newQueue = append(newQueue, k)
						}
					}
				}
			}
			queue = newQueue
		}
	}
	return result
}
