/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
    postNoRoot := len(postorder) - 1
    
    var inOrder  func(preStart, preEnd, postStart int, pre, post []int )  *TreeNode
    inOrder = func(preStart, preEnd, postStart int, pre, post []int )  *TreeNode {
        if preStart > preEnd {return nil}

        if preStart == preEnd {
            node := &TreeNode{}
            node.Val = pre[preStart]
            return node
        }

        leftRoot := pre[preStart+1]
        noLeftNode := 1
        for post[postStart + noLeftNode - 1] != leftRoot {
            noLeftNode++
        }

        node :=&TreeNode{}
        node.Val = pre[preStart]
        node.Left = inOrder(preStart+1, preStart+noLeftNode, postStart, pre, post)
        node.Right = inOrder(preStart + noLeftNode + 1, preEnd, postStart + noLeftNode , pre, post )
        return node
    }   

    res := inOrder(0, postNoRoot, 0, preorder, postorder)
    return res
}