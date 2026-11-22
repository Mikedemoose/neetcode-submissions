/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
    nodeMap := make(map[*Node]*Node)

	temp := head
	for temp != nil {
		nodeMap[temp] = &Node {
			Val: temp.Val,
		}
		temp = temp.Next
	}

	temp = head
	newList := nodeMap[temp]
	newTemp := newList

	for temp != nil {

		if temp.Next != nil {
			newTemp.Next = nodeMap[temp.Next]
		}

		if temp.Random != nil {
			newTemp.Random = nodeMap[temp.Random]
		}

		temp = temp.Next
		newTemp = newTemp.Next

	}

	return newList
}
