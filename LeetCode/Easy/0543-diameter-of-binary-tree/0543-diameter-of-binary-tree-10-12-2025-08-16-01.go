/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    maxi := 0
    childPath(root,&maxi)
    return maxi
}
func childPath(root *TreeNode, maxi *int) int {
    if root == nil {
        return 0
    }
    right := childPath(root.Right,maxi)
    left := childPath(root.Left,maxi)
    *maxi = max(*maxi,right+left)
    return max(right+1,left+1)
}