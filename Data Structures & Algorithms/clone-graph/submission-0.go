/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node == nil {
		return nil
	}

	copiedNode := &Node {
		Val: node.Val,
		Neighbors: make([]*Node, 0),
	}

	queue, copyQueue := []*Node{node}, []*Node{copiedNode}
	nodeValMap := map[int]*Node {
		node.Val: copiedNode,
	}

	visited := make(map[int]struct{})


	for len(queue) > 0 {
		curr, currCopy := queue[0], copyQueue[0]
		queue, copyQueue = queue[1:], copyQueue[1:]

		if _, ok := visited[curr.Val]; ok {
			continue
		}
		visited[curr.Val] = struct{}{}

		for _, item := range curr.Neighbors {
			var newItem *Node
			if _, ok := nodeValMap[item.Val]; ok {
				newItem = nodeValMap[item.Val]
			} else {
				newItem = &Node {
					Val: item.Val,
					Neighbors: make([]*Node, 0),
				}
				nodeValMap[item.Val] = newItem
			}
			currCopy.Neighbors = append(currCopy.Neighbors, newItem)

			if _, ok := visited[item.Val]; !ok {
				queue, copyQueue = append(queue, item), append(copyQueue, newItem)
			}
		}
	}

	

	return copiedNode
}
