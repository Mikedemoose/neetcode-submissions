/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(root *TreeNode, min, max int64) bool {
	if (root == nil) || (root.Left == nil && root.Right == nil && int64(root.Val) > min && int64(root.Val) < max) {
		return true
	} else if (root.Left != nil && root.Left.Val >= root.Val) || (root.Right != nil && root.Right.Val <= root.Val) || (int64(root.Val) >= max || int64(root.Val) <= min) {
		return false
	}
	return dfs(root.Left, min, int64(root.Val)) && dfs(root.Right, int64(root.Val), max)
}

func isValidBST(root *TreeNode) bool {
	maxVal := int64(1e9 + 1)

	return dfs(root, -maxVal, maxVal)
}
