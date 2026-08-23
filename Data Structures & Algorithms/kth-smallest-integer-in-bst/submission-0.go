/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(root *TreeNode, inorderArray *[]int, k int) {
	if root == nil {
		return
	}
	dfs(root.Left, inorderArray, k)
	*inorderArray = append(*inorderArray, root.Val)
	if len(*inorderArray) >= k {
		return
	}
	dfs(root.Right, inorderArray, k)
}

func kthSmallest(root *TreeNode, k int) int {
    inorderArray := make([]int, 0)

	dfs(root, &inorderArray, k)

	return inorderArray[k-1]
}
