/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    return height(root) != -1
}
func height(root *TreeNode) int {
    if root == nil {
        return 0
    }
    rh := height(root.Right)
    if rh == -1 {return -1}
    lh := height(root.Left)
    if lh == -1 {return -1}
    if abs(lh-rh) > 1 {return -1}
    return max(lh,rh)+1
}
func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}