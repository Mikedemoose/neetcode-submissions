func findRedundantConnection(edges [][]int) []int {
    // saves the root parent node for each node (lowest node number which can represent the subgraph)
	parentNodes := make(map[int]int)

	res := make([]int, 0)

	for _, edge := range edges {
		if isSameSubgraph(parentNodes, edge) {
			res = edge
		}
		// fmt.Println("Current state of parent Node map", parentNodes)
	}

	return res
}


func isSameSubgraph(parentNodes map[int]int, edge []int) bool {

	// fmt.Println("checking if edge", edge, "has nodes from the same subgraph")
	leftPar := findAndSetParent(parentNodes, edge[0])
	rightPar := findAndSetParent(parentNodes, edge[1])
	// fmt.Println("leftPar", leftPar, "rightPar", rightPar)

	if leftPar == rightPar {
		// fmt.Println("Edge is from the same subgraph")
		return true
	}

	newParent := min(leftPar, rightPar)
	// fmt.Println("setting new parent for both edges as", newParent)
	parentNodes[edge[0]], parentNodes[edge[1]], parentNodes[leftPar], parentNodes[rightPar] = newParent, newParent, newParent, newParent
	return false

}

func findAndSetParent(parentNodes map[int]int, node int) int {
	temp := node
	parentNode, ok := parentNodes[node]
	if !ok {
		parentNodes[node] = node
		return node
	}

	for parentNode != temp {
		temp = parentNode
		parentNode, _ = parentNodes[parentNode]
	}

	parentNodes[node] = parentNode
	return parentNode
}

