/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    fast := head.Next
    slow := head
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    h2 := sortList(slow.Next)
    slow.Next = nil
    h1 := sortList(head)
    return merge(h1,h2)
}
func merge(h1 *ListNode,h2 *ListNode) *ListNode {
    d := &ListNode{0,nil}
    tmp := d
    for h1 != nil && h2 != nil {
        if h1.Val <= h2.Val {
           tmp.Next = h1
           h1 = h1.Next
           tmp = tmp.Next
        }else{
           tmp.Next = h2
           h2 = h2.Next
           tmp = tmp.Next
        }
    }
    if h1 != nil {
        tmp.Next = h1
    }
    if h2 != nil {
        tmp.Next = h2
    }
    return d.Next
}

// //  SC O(n) array extraspace
// func sortList(head *ListNode) *ListNode {
//     tmp := head 
//     c := 0
//     for tmp != nil {
//         c++
//         tmp = tmp.Next
//     }
//     ints := make([]int,c)
//     tmp = head
//     c = 0
//     for tmp != nil {
//         ints[c] = tmp.Val
//         c++
//         tmp = tmp.Next
//     }
//     sort.Ints(ints)
//     tmp = head
//     c = 0
//     for tmp != nil {
//         tmp.Val = ints[c] 
//         c++
//         tmp = tmp.Next
//     }
//     return head
// }