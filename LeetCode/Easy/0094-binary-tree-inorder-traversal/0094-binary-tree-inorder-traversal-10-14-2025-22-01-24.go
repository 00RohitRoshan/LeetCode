/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  morris inroder traversal
func inorderTraversal(root *TreeNode) []int {
    var result []int
    current := root
    for current != nil {
        if current.Left == nil {
            result = append(result,current.Val)
            current = current.Right
        }else{
            prev := current.Left
            for prev.Right != nil && prev.Right != current {
                prev = prev.Right
            }
            if prev.Right == nil {
                prev.Right = current
                current = current.Left
            }
            if prev.Right == current {
                result = append(result,current.Val)
                current = current.Right
                prev.Right = nil
            }
        }
    }
    return result
}

// // a lot of memmory allocation
// func inorderTraversal(root *TreeNode) []int {
//     if root == nil{
//         return []int{}
//     }
//     left := inorderTraversal(root.Left)
//     right := inorderTraversal(root.Right)
//     left = append(left,root.Val)
//     left = append(left,right...)
//     return left
// }