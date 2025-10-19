/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    tmp := head 
    c := 0
    for tmp != nil {
        c++
        tmp = tmp.Next
    }
    ints := make([]int,c)
    tmp = head
    c = 0
    for tmp != nil {
        ints[c] = tmp.Val
        c++
        tmp = tmp.Next
    }
    sort.Ints(ints)
    tmp = head
    c = 0
    for tmp != nil {
        tmp.Val = ints[c] 
        c++
        tmp = tmp.Next
    }
    return head
}