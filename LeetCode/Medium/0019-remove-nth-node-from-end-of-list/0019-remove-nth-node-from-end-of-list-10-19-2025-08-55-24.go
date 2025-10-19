/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    fake := &ListNode{
        0,head,
    }
    tmp := fake
    for i:=0; i<n; i++{
        tmp = tmp.Next
    }
    source := fake
    for tmp.Next != nil {
        source = source.Next
        tmp = tmp.Next
    }
    source.Next = source.Next.Next
    return fake.Next
}