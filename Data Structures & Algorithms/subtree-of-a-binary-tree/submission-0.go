/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {

	if subRoot == nil {
		return true
	}
	if root == nil {
		return false
	}

	return isExactMatchPresent(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isExactMatchPresent(root, subRoot *TreeNode) bool {
	if subRoot == nil && root == nil {
		return true
	} else if root == nil || subRoot == nil {
		return false
	}

	if root.Val != subRoot.Val {
		return false
	}

	return isExactMatchPresent(root.Left, subRoot.Left) && isExactMatchPresent(root.Right, subRoot.Right)
}
