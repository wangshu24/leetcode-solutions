/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
    branchLen := len(preorder)-1
    var inOrder func(preStart, preEnd, postStart int) *TreeNode 
    inOrder = func(preStart, preEnd, postStart int) *TreeNode {
        if preStart > preEnd { return nil }

        if preStart == preEnd {
            node := &TreeNode{}
            node.Val = preorder[preStart]
            return node
        }

        leftRoot := preorder[preStart+1]
        noLeftNode := 1
        for postorder[postStart + noLeftNode -1] != leftRoot {
            noLeftNode++
        }

        node := &TreeNode{}
        node.Val = preorder[preStart]
        node.Left = inOrder(preStart+1, preStart + noLeftNode, postStart)
        node.Right = inOrder(preStart + noLeftNode+1, preEnd, postStart + noLeftNode)
        return node
    }

    res := inOrder(0, branchLen, 0)
    return res
}