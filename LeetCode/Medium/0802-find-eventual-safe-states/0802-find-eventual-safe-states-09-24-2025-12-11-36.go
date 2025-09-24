func eventualSafeNodes(graph [][]int) []int {
    n := len(graph)
    reverseGraph := make([][]int,n)
    inDegree := make([]int,n)
    for node,adj := range graph {
        for _,newNode := range adj {
            reverseGraph[newNode] = append(reverseGraph[newNode],node)
            inDegree[node]++
        }
    }
    leaves := []int{}
    for leaf,in := range inDegree {
        if in == 0 {
            leaves = append(leaves,leaf)
        }
    }
    ans := []int{}
    for len(leaves) > 0 {
        newLeaves := []int{}
        for _,leaf := range leaves {
            ans = append(ans,leaf)
            for _,adj := range reverseGraph[leaf] {
                inDegree[adj]--
                if inDegree[adj] == 0 {
                    newLeaves = append(newLeaves,adj)
                }
            }
        }
        leaves = newLeaves
    }
    sort.Ints(ans)
    return ans
}