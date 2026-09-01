/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    
	nodeToCopyMap := make(map[*Node]*Node)

	dummyHead := &Node {}

	curr := head
	newCurr := dummyHead

	for curr != nil {
		var nodeFromMap *Node
		var ok bool
		if nodeFromMap, ok = nodeToCopyMap[curr]; !ok {
			nodeFromMap = &Node{
				Val: curr.Val,
			}
			nodeToCopyMap[curr] = nodeFromMap
		}

		next, random := curr.Next, curr.Random
		if next != nil {
			var newNext *Node
			if newNext, ok = nodeToCopyMap[next]; !ok {
				newNext = &Node {
					Val: next.Val,
				}
				nodeToCopyMap[next] = newNext
			}
			nodeFromMap.Next = newNext
		}
		if random != nil {
			var newrandom *Node
			if newrandom, ok = nodeToCopyMap[random]; !ok {
				newrandom = &Node {
					Val: random.Val,
				}
				nodeToCopyMap[random] = newrandom
			}
			nodeFromMap.Random = newrandom
		}
		newCurr.Next = nodeFromMap
		newCurr = newCurr.Next
		curr = curr.Next

	}
	return dummyHead.Next
}
