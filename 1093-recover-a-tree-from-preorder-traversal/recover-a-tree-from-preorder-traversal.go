/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverFromPreorder(traversal string) *TreeNode {
    // prv[i]: last node created at level i
    prv := make(map[int]*TreeNode)
    // create the root node (always present)
    _,v := next(&traversal)
    prv[0] = &TreeNode{v, nil, nil}
    for len(traversal) > 0 {
        // create the next node at level L
        l,v := next(&traversal)
        prv[l] = &TreeNode{v, nil, nil}
        // link the new node to the last parent at level L-1
        if p:=prv[l-1]; (*p).Left == nil {
            // create left child if available
            (*p).Left = prv[l] 
        } else { 
            // otherwise create right child
            (*p).Right = prv[l] 
        }
    }
    // return the root
    return prv[0]
}

// read the next level and value from the traversal string
func next(t *string) (int, int) {
    p, lvl, val := 0, 0, 0
    // number of dashes indicate the level
    for (*t)[p] == '-' { p, lvl = p+1, lvl+1 }
    // read digits and calculate their value
    for p<len(*t) && (*t)[p] != '-' { 
        p, val = p+1, val*10+int((*t)[p]-'0') 
    }
    // remove the characters already read
    *t = (*t)[p:]
    return lvl, val
}