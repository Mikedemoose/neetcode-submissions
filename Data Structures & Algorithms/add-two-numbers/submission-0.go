/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    sum := &ListNode{}
	sumPtr := sum

	queue := []*ListNode {l1, l2}
	carry := 0

	for len(queue) > 0 {
		newQueue := make([]*ListNode, 0)

		currSum := 0

		for _, n := range queue{
			currSum += n.Val
			if n.Next != nil {
				newQueue = append(newQueue, n.Next)
			}
		}
		currSum += carry

		sumPtr.Next = &ListNode {
			Val: currSum%10,
		}
		sumPtr = sumPtr.Next

		carry = currSum/10
		queue = newQueue
	}

	if carry != 0 {
		sumPtr.Next = &ListNode{
			Val: carry,
		}
	}

	return sum.Next
}
