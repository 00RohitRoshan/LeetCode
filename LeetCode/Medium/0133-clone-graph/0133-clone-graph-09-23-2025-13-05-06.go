/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
func cloneGraph(node *Node) *Node {
    vis := map[int]*Node{}
    return Clone(node,vis)
}
func Clone(node *Node,vis map[int]*Node) *Node {
    if node == nil {
        return nil
    }
    if vis == nil {
        vis = map[int]*Node{}
    }
    if vis[node.Val] != nil {
        return vis[node.Val]
    }
    n := &Node{
        Val:       node.Val,
        Neighbors: []*Node{},
    }
    // \U0001f527 Store current node before recursion to avoid cycles
    vis[node.Val] = n
    for _, v := range node.Neighbors {
        n.Neighbors = append(n.Neighbors, Clone(v,vis))
    }
    return n
}
