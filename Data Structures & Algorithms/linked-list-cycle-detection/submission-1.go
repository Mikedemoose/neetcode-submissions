/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    fast, slow := head, head

	for fast != nil && slow != nil {
		fast = fast.Next
		slow = slow.Next
		if slow != nil && slow.Next != nil {
			slow = slow.Next
		} else {
			return false
		}

		if fast == slow {
			return true
		}
	}

	return false
}
