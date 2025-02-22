/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type FindElements struct {
    root *TreeNode
}


func Constructor(root *TreeNode) FindElements {
    // simple store the root 
    res := FindElements{root: root}
    return res
}


func (this *FindElements) Find(target int) bool {
    node := this.root
    // calculate the mask to iterate over target bits 
    // from the left to the right
    mask := 1
    num := target + 1
    for num != 0 {
        num >>= 1
        mask <<= 1
    }
    mask >>= 1
    num = (target + 1) & (mask - 1)
    mask >>= 1
    for mask != 0 {
        if num & mask == 0 {
            if node.Left == nil {
                return false
            }
            node = node.Left 
        } else {
            if node.Right == nil {
                return false
            }
            node = node.Right
        }
        mask >>= 1
    }
    return true
}


/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */