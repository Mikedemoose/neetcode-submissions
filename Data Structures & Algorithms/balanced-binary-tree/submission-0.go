/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    val, _ := recurse(root)
	return val
}

func recurse(root *TreeNode) (bool, int) {

	if root == nil {
		return true, 0
	}

	var leftHeight, rightHeight int
	leftok, rightok := true, true

	if root.Left != nil {
		leftok, leftHeight = recurse(root.Left)
		if !leftok {
			return false, 0
		}
		leftHeight++
	}
	if root.Right != nil {
		rightok, rightHeight = recurse(root.Right)
		if !rightok {
			return false, 0
		}
		rightHeight++
	}

	if rightHeight < leftHeight {
		rightHeight, leftHeight = leftHeight, rightHeight
	}

	if rightHeight-leftHeight > 1 {
		return false, 0
	}

	return true, rightHeight
}
