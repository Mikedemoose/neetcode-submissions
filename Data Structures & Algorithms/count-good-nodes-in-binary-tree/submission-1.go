/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    return getNumGoodNodes(root, root.Val)
}

func getNumGoodNodes(root *TreeNode, currMax int) int {
	if root == nil {
		return 0
	}
	total := 0
	if root.Val >= currMax {
		total += 1
	}

	return total + getNumGoodNodes(root.Left, max(currMax, root.Val)) + getNumGoodNodes(root.Right, max(currMax, root.Val))
}
