/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func reverseOddLevels(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	queue := []*TreeNode{root}
	level := 0

	for len(queue) > 0 {
		n := len(queue)

		// if level is odd
		if level%2 == 1 {
			for i, j := 0, len(queue)-1; i < j; i, j = i+1, j-1 {
				queue[i].Val, queue[j].Val = queue[j].Val, queue[i].Val
			}
		}

		// append the next level
		// with respect to queue
		for i := range n {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		queue = queue[n:]
		level++
	}

	return root
}