/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var depthSum []int

func dfs1(root *TreeNode, d int) {
    if root == nil {
        return 
    }

    if d >= len(depthSum) {
        depthSum = append(depthSum, root.Val)
    } else {
        depthSum[d] += root.Val
    }

    dfs1(root.Left, d+1)
    dfs1(root.Right, d+1)
}

func dfs2(root *TreeNode, val int, d int) {
    if root == nil {
        return 
    }

    root.Val = val

    c := 0
    if d+1 < len(depthSum) {
        c = depthSum[d+1]
    }
    
    if root.Left != nil {
        c -= root.Left.Val
    }
    if root.Right != nil {
        c -= root.Right.Val
    }

    if root.Left != nil {
        dfs2(root.Left, c, d+1)
    }
    if root.Right != nil {
        dfs2(root.Right, c, d+1)
    }
}

func replaceValueInTree(root *TreeNode) *TreeNode {
    depthSum = []int{} 
    dfs1(root, 0)
    dfs2(root, 0, 0)
    return root
}