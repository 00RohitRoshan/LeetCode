/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    return isSame(root.Left,root.Right)
}
func isSame(r1 *TreeNode,r2 *TreeNode) bool {
    if r1 == nil || r2 == nil {
        return r1 == nil && r2 == nil
    }
    if isSame(r1.Right,r2.Left) && isSame(r2.Right,r1.Left){
        return r1.Val == r2.Val
    }
    return false
}