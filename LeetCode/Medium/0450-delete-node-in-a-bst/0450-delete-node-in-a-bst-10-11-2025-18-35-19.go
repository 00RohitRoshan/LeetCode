/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  instead of checking once for nil reach the node until it becomes nil
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return nil
    }
    if root.Val == key {
        if root.Left != nil {
            tmpRoot := root.Left
            for tmpRoot.Right != nil {
                tmpRoot = tmpRoot.Right
            }
            tmpRoot.Right = root.Right
            return root.Left
        }
        if root.Right != nil {
            tmpRoot := root.Right
            for tmpRoot.Left != nil {
                tmpRoot = tmpRoot.Left
            }
            tmpRoot.Left = root.Left
            return root.Right
        }
        return nil
    }
    root.Left = deleteNode(root.Left,key)
    root.Right = deleteNode(root.Right,key)
    return root
}

// //  this is wrong as it is a BST we can't attach large node to left and small node to right
// func deleteNode(root *TreeNode, key int) *TreeNode {
//     if root == nil {
//         return nil
//     }
//     if root.Val == key {
//         if root.Left != nil {
//             if root.Left.Right == nil {
//                 root.Left.Right = root.Right
//                 return root.Left
//             }
//             if root.Left.Left == nil {
//                 root.Left.Left = root.Right
//                 return root.Left
//             }
//         }
//         if root.Right != nil {
//             if root.Right.Left == nil {
//                 root.Right.Left = root.Left
//                 return root.Right
//             }
//             if root.Right.Right == nil {
//                 root.Right.Right = root.Left
//                 return root.Right
//             }
//         }
//         return nil
//     }
//     root.Left = deleteNode(root.Left,key)
//     root.Right = deleteNode(root.Right,key)
//     return root
// }