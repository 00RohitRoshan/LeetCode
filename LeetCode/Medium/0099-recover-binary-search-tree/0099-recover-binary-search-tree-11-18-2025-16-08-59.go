/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode) {
	var prevNode, a, b *TreeNode
	inOrder(root, func(p *TreeNode) {
		if prevNode != nil && prevNode.Val > p.Val {
			if a == nil {
				a = prevNode
			}
			b = p
		}
		prevNode = p
	})

	swap(a, b)
}
func inOrder(root *TreeNode, doFunc func(p *TreeNode)) {
	if root == nil {
		return
	}
	inOrder(root.Left, doFunc)
	doFunc(root)
	inOrder(root.Right, doFunc)
}
func swap(p, q *TreeNode) {
	if p != nil && q != nil {
		p.Val, q.Val = q.Val, p.Val
	}
}