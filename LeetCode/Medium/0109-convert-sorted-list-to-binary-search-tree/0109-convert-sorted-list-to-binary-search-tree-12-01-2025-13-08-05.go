/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil {
        return nil
    }
    if head.Next == nil {
        return &TreeNode{Val: head.Val}
    }
    fast := head
    slow := head
    var prev *ListNode
    for fast != nil && fast.Next != nil {
        prev = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    if prev != nil {
        prev.Next = nil
    }
    root := &TreeNode{slow.Val,nil,nil}
    root.Right = sortedListToBST(slow.Next)
    root.Left = sortedListToBST(head)
    return root
}

// func insert(root *TreeNode,v int){
//     if v > root.Val {
//         if root.Right != nil {
//             insert(root.Right,v)
//         }else{
//             root.Right = &TreeNode{v,nil,nil}
//         }
//     }else{
//         if root.Left != nil {
//             insert(root.Left,v)
//         }else{
//             root.Left = &TreeNode{v,nil,nil}
//         }
//     }
// }

// func reverse(head *ListNode) *ListNode {
//     prev := nil 
//     cur := head 
//     for cur != nil {
//         next := cur.Next
//         cur.Next = prev
//         prev = cur
//         cur = next
//     }
//     return prev
// }