/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func getOrNewItem(val int, m map[int]*Node) *Node {
	if val, ok := m[val]; ok {
		return val
	}
	newItem := &Node {
		Val: val,
		Neighbors: make([]*Node, 0),
	}
	m[val] = newItem
	return newItem
}

func cloneGraph(node *Node) *Node {
    if node == nil {
		return nil
	}

	copiedNode := &Node {
		Val: node.Val,
		Neighbors: make([]*Node, 0),
	}

	nodeValMap := map[int]*Node{
		node.Val: copiedNode,
	}

	visited := make(map[int]struct{})

	queue := []*Node{node}

	for len(queue) > 0 {

		curr := queue[0]
		queue = queue[1:]

		if _, ok := visited[curr.Val]; ok {
			continue
		}
		visited[curr.Val] = struct{}{}

		newItem := getOrNewItem(curr.Val, nodeValMap)
		for _, item := range curr.Neighbors {
			newItem.Neighbors = append(newItem.Neighbors, getOrNewItem(item.Val, nodeValMap))
			if _, ok := visited[item.Val]; !ok {
				queue = append(queue, item)
			}
		}

	}

	

	return copiedNode
}
