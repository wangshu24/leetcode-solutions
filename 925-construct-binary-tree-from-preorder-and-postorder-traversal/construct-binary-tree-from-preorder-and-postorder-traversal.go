type queueItem struct {
	Node *TreeNode
	Pre  []int
	Post []int
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	var queue []queueItem

	root := &TreeNode{Val: preorder[0]}
	preorder = preorder[1:]
	postorder = postorder[:len(postorder)-1]

	if len(preorder) < 1 {
		return root
	}

	x := 0
	for preorder[x] != postorder[len(postorder)-1] {
		x++
	}

	preLeftTree := preorder[:x]
	postLeftTree := postorder[:x]
	preRightTree := preorder[x:]
	postRightTree := postorder[x:]

	if len(preLeftTree) > 0 {
		root.Left = &TreeNode{Val: preLeftTree[0]}
		preLeftTree = preLeftTree[1:]
		postLeftTree = postLeftTree[:len(postLeftTree)-1]
		if len(preLeftTree) > 0 {
			queue = append(queue, queueItem{Node: root.Left, Pre: preLeftTree, Post: postLeftTree})
		}
	}
	if len(preRightTree) > 0 {
		root.Right = &TreeNode{Val: preRightTree[0]}
		preRightTree = preRightTree[1:]
		postRightTree = postRightTree[:len(postRightTree)-1]
		if len(preRightTree) > 0 {
			queue = append(queue, queueItem{Node: root.Right, Pre: preRightTree, Post: postRightTree})
		}
	}
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		x = 0
		for item.Pre[x] != item.Post[len(item.Post)-1] {
			x++
		}

		preLeftTree = item.Pre[:x]
		postLeftTree = item.Post[:x]
		preRightTree = item.Pre[x:]
		postRightTree = item.Post[x:]

		if len(preLeftTree) > 0 {
			item.Node.Left = &TreeNode{Val: preLeftTree[0]}
			preLeftTree = preLeftTree[1:]
			postLeftTree = postLeftTree[:len(postLeftTree)-1]
			if len(preLeftTree) > 0 {
				queue = append(queue, queueItem{Node: item.Node.Left, Pre: preLeftTree, Post: postLeftTree})
			}
		}
		if len(preRightTree) > 0 {
			item.Node.Right = &TreeNode{Val: preRightTree[0]}
			preRightTree = preRightTree[1:]
			postRightTree = postRightTree[:len(postRightTree)-1]
			if len(preRightTree) > 0 {
				queue = append(queue, queueItem{Node: item.Node.Right, Pre: preRightTree, Post: postRightTree})
			}
		}
	}
	return root
}