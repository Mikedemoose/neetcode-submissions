/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) <= 0 || len(inorder) <= 0 {
		return nil
	}
    // find the position of the root element in inorder array
	rootVal := preorder[0]
	rootPos := 0

	for i, val := range inorder {
		if val == rootVal {
			rootPos = i
			break
		}
	}

	root := &TreeNode {
		Val: rootVal,
	}

	root.Left = buildTree(preorder[1:rootPos+1], inorder[:rootPos])
	root.Right = buildTree(preorder[rootPos+1:], inorder[rootPos+1:])
	return root
}
