/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var prev *ListNode
	curr, next := head, head.Next

	for next != nil {
		curr.Next = prev
		prev = curr
		curr = next
		next = next.Next
		curr.Next = prev
	}

	return curr
}
