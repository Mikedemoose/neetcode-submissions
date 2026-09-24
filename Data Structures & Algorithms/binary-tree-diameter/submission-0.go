/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    diameter := 0
	recurse(root, &diameter)

	return diameter
}

func recurse(root *TreeNode, diameter *int) int {
	if root == nil {
		return 0
	}

	var longestLeft, longestRight int

	if root.Left != nil {
		longestLeft = recurse(root.Left, diameter) + 1
	}
	if root.Right != nil {
		longestRight = recurse(root.Right, diameter) + 1
	}

	*diameter = max(*diameter, longestLeft + longestRight)
	return max(longestLeft, longestRight)
}
