/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	maxHeight := 0

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		maxHeight++

		newQueue := make([]*TreeNode, 0)

		for _, item := range queue {
			if item != nil {
				if item.Left != nil {
					newQueue = append(newQueue, item.Left)
				}
				if item.Right != nil {
					newQueue = append(newQueue, item.Right)
				}
			}
		}
		queue = newQueue
	}

	return maxHeight
    
}
