func isBipartite(graph [][]int) bool {
    visited := make([]int, len(graph))
    for i := range graph {
        if visited[i] == 0{
            q := []int{i}
            visited[i] = 1
            for len(q) > 0{
                cur := q[0]
                q = q[1:]
                for _,k := range graph[cur] {
                    if visited[k] == visited[cur] {
                        return false
                    }
                    if visited[k] == 0{
                        visited[k] = -visited[cur]
                        q = append(q,k)
                    }
                }
            }
        }
    }
    return true
}