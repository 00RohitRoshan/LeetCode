func findCircleNum(isConnected [][]int) int {
    n := len(isConnected)
    visited := map[int]bool{}
    count := 0
    q := []int{}
    for i := range n {
        if !visited[i] {
            count++
            visited[i] = true
            q = append(q, i)

        }
            for len(q) > 0 {
                cur := q[0]
                q = q[1:] // dequeue

                for j := range n {
                    if isConnected[cur][j] == 1 && !visited[j] {
                        visited[j] = true
                        q = append(q, j)
                    }
                }
            }
    }
    return count
}