/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
    currQueue := []*TreeNode{root}
	finalList := make([]int, 0)

	for len(currQueue) > 0 {
		finalList = append(finalList, currQueue[len(currQueue)-1].Val)
		newQueue := make([]*TreeNode, 0)

		for _, item := range currQueue {
			if item.Left != nil {
				newQueue = append(newQueue, item.Left)
			}
			if item.Right != nil {
				newQueue = append(newQueue, item.Right)
			}
		}

		currQueue = newQueue
	}

	return finalList
}
