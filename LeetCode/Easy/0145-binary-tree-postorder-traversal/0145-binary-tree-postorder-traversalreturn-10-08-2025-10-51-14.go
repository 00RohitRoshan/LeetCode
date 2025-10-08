/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func postorderTraversal(root *TreeNode) []int {
//     res := []int{}
//     traverse(root,&res)
//     return res
// }
// func traverse(root *TreeNode,arr *[]int){
//     if root == nil {
//         return
//     }
//     traverse(root.Left,arr)
//     traverse(root.Right,arr)
//     *arr = append(*arr,root.Val)
// }

func postorderTraversal(root *TreeNode) []int {
    return traverse(root)
}
func traverse(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    left := traverse(root.Left)
    right := traverse(root.Right)
    return append(append(left, right...), root.Val)
}
