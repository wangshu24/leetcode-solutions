/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
   xl, yl := dfs(root, x, 0 ), dfs(root, y, 0)
    fmt.Println("depth for x: ", xl, " and dep for y: ", yl)
    isBro := isBrother(root,x, y)
    fmt.Println(isBro)
   return xl == yl && !isBro
}

func dfs(root *TreeNode, a,dep int) int {
    if root == nil {return 0}
    dep++
    if root.Val == a {return dep}

    left := dfs(root.Left, a, dep)
    right := dfs(root.Right, a, dep)

    return max(left, right) 
}

func isBrother(root *TreeNode, a, b int) bool {
    if root == nil {return false}
    if root.Left == nil {
        return isBrother(root.Right, a, b)
    } 
    if root.Right == nil {
        return isBrother(root.Left, a, b)
    }
    
    if ((root.Left.Val == a) && (root.Right.Val == b)) || ((root.Left.Val == b) && (root.Right.Val == a)) {return true} else {
        return isBrother(root.Left, a, b) || isBrother(root.Right, a,b)
    }
}

func max(a, b int) int{
    if a > b {return a} 
    return b
}