/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
    currentQueue := make([]*TreeNode, 0)

	currentQueue = append(currentQueue, root)
	finalList := make([][]int, 0)
	finalList = append(finalList, []int{root.Val})

	for len(currentQueue) > 0 {
		childQueue := make([]*TreeNode, 0)
		childList := make([]int, 0)

		for _, curr := range currentQueue {
			if curr.Left != nil {
				childQueue = append(childQueue, curr.Left)
				childList = append(childList, curr.Left.Val)
			}
			if curr.Right != nil {
				childQueue = append(childQueue, curr.Right)
				childList = append(childList, curr.Right.Val)
			}

		}

		currentQueue = childQueue
		if len(childList) > 0 {
			finalList = append(finalList, childList)
		}
	}

	return finalList
}
