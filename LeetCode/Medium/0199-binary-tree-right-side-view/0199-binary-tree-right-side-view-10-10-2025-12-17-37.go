/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    res := []int{}
    q := []*TreeNode{root}
    for len(q) > 0 {
        nextq := []*TreeNode{}
        for i,node := range q {
            if node != nil {
                if node.Left != nil && node.Right == nil{
                    nextq = append(nextq,node.Left)
                }else if node.Left == nil && node.Right != nil{
                    nextq = append(nextq,node.Right)
                }else if node.Left != nil && node.Right != nil{
                    nextq = append(nextq,node.Left,node.Right)
                }
                if i == len(q)-1 {
                    res = append(res,node.Val)
                }
            }
        }
        q = nextq
    }
    return res
}