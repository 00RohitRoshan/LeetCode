/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }
    vis := map[int]*Node{}
    vis[node.Val] = &Node{
        Val : node.Val,
        Neighbors : []*Node{},
    }
    q := []*Node{node}
    for len(q) > 0 {
        cur := q[0]
        q = q[1:]
        for _,v := range cur.Neighbors {
            if vis[v.Val] == nil {
                vis[v.Val] = &Node{
                    Val : v.Val,
                    Neighbors : []*Node{},
                }
                q = append(q,v)
            }
            vis[cur.Val].Neighbors = append(vis[cur.Val].Neighbors,vis[v.Val])
        }
    }
    return vis[node.Val]
}
// func Clone(node *Node,vis map[int]*Node) *Node {
//     if node == nil {
//         return nil
//     }
//     if vis == nil {
//         vis = map[int]*Node{}
//     }
//     if vis[node.Val] != nil {
//         return vis[node.Val]
//     }
//     n := &Node{
//         Val:       node.Val,
//         Neighbors: []*Node{},
//     }
//     // \U0001f527 Store current node before recursion to avoid cycles
//     vis[node.Val] = n
//     for _, v := range node.Neighbors {
//         n.Neighbors = append(n.Neighbors, Clone(v,vis))
//     }
//     return n
// }
