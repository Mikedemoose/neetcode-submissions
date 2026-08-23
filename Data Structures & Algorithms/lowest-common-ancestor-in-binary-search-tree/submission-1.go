/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	tempRoot := root

	if q.Val < p.Val {
		temp := p
		p = q
		q = temp
	}

	for tempRoot != nil {
		if q.Val < tempRoot.Val {
			tempRoot = tempRoot.Left
		} else if p.Val > tempRoot.Val {
			tempRoot = tempRoot.Right
		} else {
			return tempRoot
		}
	}

	return nil
}
