/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    small := make([]int, 0)

    var bst func(*TreeNode)
    bst = func(tmp *TreeNode) {
        if tmp == nil {
            return
        }
        bst(tmp.Left)
        small = append(small, tmp.Val)
        bst(tmp.Right)
        return
    }
    bst(root)
    fmt.Println(small)

    // Now that we have all the node in order, retrieved via in-order traversal
    // We can easily count from smallest to get kth smallest value
    return small[k-1]
}

