/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    return getGoodNodeCountDfs(root, root.Val)
}

func getGoodNodeCountDfs(root *TreeNode, currMax int) int {
	if root == nil {
		return 0
	}

	count := 0

	if root.Val >= currMax {
		count++
		currMax = root.Val
	}

	return count + getGoodNodeCountDfs(root.Left, currMax) + getGoodNodeCountDfs(root.Right, currMax)
}
