/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    rootIndex := map[int]int{}
    for i,v := range inorder {
        rootIndex[v] = i
    }
    i := 0
    var find func(int,int)*TreeNode
    find = func (j int,k int) *TreeNode {
        if j > k {
            return nil
        }
        nodeVal := preorder[i]
        i++
        splitIdx := rootIndex[nodeVal]
        node := &TreeNode{Val:nodeVal}
        node.Left = find(j,splitIdx-1)
        node.Right = find(splitIdx+1,k)
        return node
    }
    return find(0,len(inorder)-1)
}
