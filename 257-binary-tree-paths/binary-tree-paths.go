/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func binaryTreePaths(root *TreeNode) []string {
    var res []string

    var dfs func (root *TreeNode, str string)
    dfs = func(root *TreeNode, str string) {
        if root == nil {return}
        str+=strconv.Itoa(root.Val)
        if root.Left == nil && root.Right == nil {
            res = append(res, str)
        }

        str+="->"
        dfs(root.Left, str)
        dfs(root.Right, str)
    }
    dfs(root, "")

    return res
}

// func dfs(root *TreeNode, str string, res []string)  {
//     if root == nil {return}
//     str+= strconv.Itoa(root.Val)
//     if root.Left == nil && root.Right == nil {
//         res = append(res, str)
//         fmt.Println(res)
//     }

//     str+="->"
//     dfs(root.Left, str, res)
//     dfs(root.Right, str, res)

//     // return res
// }
