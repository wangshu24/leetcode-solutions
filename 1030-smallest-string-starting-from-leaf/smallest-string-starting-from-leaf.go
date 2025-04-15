/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func smallestFromLeaf(root *TreeNode) string {
    var (
        res string
    )

    var dfs func(root *TreeNode, set []uint8)

    dfs = func(root *TreeNode, set []uint8) {
        if root == nil {
            return
        }

        set = append([]uint8{uint8(root.Val+97)} , set...)
        if root.Left == nil && root.Right == nil {
            if res == "" || string(set) < res{
                res = string(set)
            }
        }

        dfs(root.Left, set)
        dfs(root.Right, set)
    }

    dfs(root, []uint8{})
    return res
}   
