/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    
	// find the middle of the list

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// mid is slow

	// save the rest of the list from mid in an array
	secondHalf := make([]*ListNode, 0)
	temp := slow.Next
	for temp != nil {
		secondHalf = append(secondHalf, temp)
		temp = temp.Next
	}

	slow.Next = nil

	//loop through second half and insert it in the list
	newHead := head
	for i:=len(secondHalf)-1; i>=0; i-- {
		temp = newHead.Next
		newHead.Next = secondHalf[i]
		if newHead.Next != nil {
			newHead.Next.Next = temp
		}
		newHead = temp
	}
}
