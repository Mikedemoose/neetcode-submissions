/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func makeCopy(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	return &TreeNode{
		Val: node.Val,
	}
}

func invertTree(root *TreeNode) *TreeNode {
    
	if root == nil {
		return nil
	}

	newRoot := makeCopy(root)


	newParentQueue := []*TreeNode{newRoot}
	processQueue := []*TreeNode{root.Right, root.Left}

	for len(processQueue) > 0 {

		parent_queue := make([]*TreeNode, 0)
		child_queue := make([]*TreeNode, 0)

		p_i, c_i := 0, 0

		for p_i < len(newParentQueue) {
			copy_left, copy_right := makeCopy(processQueue[c_i]), makeCopy(processQueue[c_i+1])
			newParentQueue[p_i].Left = copy_left
			newParentQueue[p_i].Right = copy_right

			if copy_left != nil {
				parent_queue = append(parent_queue, copy_left)
				child_queue = append(child_queue, processQueue[c_i].Right, processQueue[c_i].Left)
			}
			if copy_right != nil {
				parent_queue = append(parent_queue, copy_right)
				child_queue = append(child_queue, processQueue[c_i+1].Right, processQueue[c_i+1].Left)
			}

			p_i+=1
			c_i+=2
		}

		newParentQueue = parent_queue
		processQueue = child_queue

	}

	return newRoot
}
