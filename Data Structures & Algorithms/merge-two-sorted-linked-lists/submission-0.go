/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    finalList := &ListNode{}
	curr := finalList

	left, right := list1, list2

	for left != nil && right != nil {
		if left.Val < right.Val {
			curr.Next = left
			left = left.Next
			curr = curr.Next
			curr.Next = nil
		} else {
			curr.Next = right
			right = right.Next
			curr = curr.Next
			curr.Next = nil
		}
	}

	if left != nil {
		curr.Next = left
	} else if right != nil {
		curr.Next = right
	}

	return finalList.Next
}
