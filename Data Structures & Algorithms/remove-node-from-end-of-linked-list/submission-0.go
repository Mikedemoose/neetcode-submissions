/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// base case
	// find length of the linked list
	temp := head
	length := 0
	for temp != nil {
		temp = temp.Next
		length++
	}
	if n==length {
		return head.Next
	} else if n==length-1 {
		head.Next = head.Next.Next
		return head
	}


    start := head
	follow := head
	for _ = range n+1 {
		start = start.Next
	}

	for start != nil {
		start = start.Next
		follow = follow.Next
	}

	follow.Next = follow.Next.Next

	return head

}
