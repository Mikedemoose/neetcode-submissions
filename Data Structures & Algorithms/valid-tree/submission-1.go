func validTree(n int, edges [][]int) bool {
	adjMap := make(map[int]map[int]struct{})
	for _, edge := range edges {
		if _, ok := adjMap[edge[0]]; !ok {
			adjMap[edge[0]] = make(map[int]struct{})
		}
		if _, ok := adjMap[edge[1]]; !ok {
			adjMap[edge[1]] = make(map[int]struct{})
		}
		adjMap[edge[0]][edge[1]] = struct{}{}
		adjMap[edge[1]][edge[0]] = struct{}{}
	}

	visited := make(map[int]struct{})

	queue := []int{0}
	visited[0] = struct{}{}

	for len(queue) > 0 {
		newQueue := make([]int, 0)
		for _, item := range queue {
			for k, _ := range adjMap[item] {
				if _, ok := visited[k]; !ok {
					newQueue = append(newQueue, k)
					visited[k] = struct{}{}
				}
			}
		}
		queue = newQueue
	}
	
	return len(visited)==n && n-len(edges) == 1
}
