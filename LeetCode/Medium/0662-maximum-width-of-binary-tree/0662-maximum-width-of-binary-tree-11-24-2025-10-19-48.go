/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type pair struct {
    node *TreeNode
    i    int // Position index in a complete binary tree
}

func widthOfBinaryTree(root *TreeNode) int {
    q := []pair{{root, 1}} // Start with root at index 1
    ans := 0

    for len(q) > 0 {
        levelSize := len(q)
        start := q[0].i     // Index of first node in level
        end := q[levelSize-1].i // Index of last node in level
        fmt.Println(q)
        ans = max(ans, end-start+1)

        for i := 0; i < levelSize; i++ {
            p := q[0]
            q = q[1:]
            node := p.node
            idx := p.i

            if node.Left != nil {
                q = append(q, pair{node.Left, 2*idx})
            }
            if node.Right != nil {
                q = append(q, pair{node.Right, 2*idx+ 1})
            }
        }
    }

    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}