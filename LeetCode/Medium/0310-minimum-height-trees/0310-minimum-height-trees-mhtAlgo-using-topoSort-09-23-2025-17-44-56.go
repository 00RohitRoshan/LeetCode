func findMinHeightTrees(n int, edges [][]int) []int {
    if n <= 2 {
        res := []int{}
        for i := range n {
            res = append(res,i)
        }
        return res
    }
    graph := map[int][]int{}
    inDegree := make([]int,n)
    for _, edge := range edges {
        graph[edge[0]] = append(graph[edge[0]],edge[1])
        inDegree[edge[0]]++
        graph[edge[1]] = append(graph[edge[1]],edge[0])
        inDegree[edge[1]]++
    }
    leaves := []int{}
    for leaf,degree := range inDegree {
        if degree == 1 {
            leaves = append(leaves,leaf)
        }
    }
    remainingnodes := n
    for remainingnodes > 2 {
        remainingnodes -= len(leaves)
        newLeaves := []int{}
        for _,leaf := range leaves {
            inDegree[leaf]--
            for _, neighbour := range graph[leaf] {
                inDegree[neighbour]--
                if inDegree[neighbour] == 1 {
                    newLeaves = append(newLeaves,neighbour)
                }
            }
        }
        leaves = newLeaves
    }
    return leaves
}
// TLE n^2 
// func findMinHeightTrees(n int, edges [][]int) []int {
//     minHeight := 9999
//     graph := map[int][]int{}
//     for _,edge := range edges {
//         graph[edge[0]] = append(graph[edge[0]],edge[1])
//         graph[edge[1]] = append(graph[edge[1]],edge[0])
//     }
//     fmt.Println(graph)
//     hNode := map[int][]int{}
//     for i := range n {
//         vis := make([]bool,n)
//         height := dfs(i,graph,vis)
//         hNode[height] = append(hNode[height],i)
//         minHeight = min(minHeight,height)
//     }
//     fmt.Println(hNode)
//     return hNode[minHeight]
// }
// func dfs(n int, graph map[int][]int, vis []bool)int{
//     maxHeight := 0
//     if !vis[n] {
//         vis[n] = true
//         for _,nextN := range graph[n] {
//             maxHeight = max(maxHeight,dfs(nextN,graph,vis))
//         }
//     }
//     return maxHeight+1
// }