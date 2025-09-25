func shortestPathBinaryMatrix(grid [][]int) int {
    n := len(grid)
    
    // Early exit if the start or end points are blocked
    if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
        return -1
    }

    // Directions: right, down-right, down, down-left, left, up-left, up, up-right
    moves := [][]int{{0,1},{1,1},{1,0},{1,-1},{0,-1},{-1,-1},{-1,0},{-1,1}}

    // BFS initialization
    q := [][]int{{0, 0}}  // Start from the top-left corner
    vis := make([][]bool, n)
    for i := 0; i < n; i++ {
        vis[i] = make([]bool, n)
    }
    vis[0][0] = true  // Mark the start as visited
    
    count := 1  // Starting point is already part of the path
    
    // Perform BFS
    for len(q) > 0 {
        levelSize := len(q)
        
        for i := 0; i < levelSize; i++ {
            x, y := q[i][0], q[i][1]
            
            // If we reached the bottom-right corner
            if x == n-1 && y == n-1 {
                return count
            }

            // Explore all 8 possible directions
            for _, move := range moves {
                r, s := x + move[0], y + move[1]
                if r >= 0 && r < n && s >= 0 && s < n && grid[r][s] == 0 && !vis[r][s] {
                    vis[r][s] = true
                    q = append(q, []int{r, s})
                }
            }
        }
        
        // Increment the count after processing the current level of BFS
        count++
        q = q[levelSize:]
    }

    return -1  // No path found
}

// func shortestPathBinaryMatrix(grid [][]int) int {
//     q := [][]int{{0,0}}
//     n := len(grid)
//     count := 0
//     moves := [][]int{{0,1},{1,1},{1,0},{1,-1},{0,-1},{-1,-1},{-1,0},{-1,1}}
//     vis := make(map[string]struct{})
//     for len(q) > 0 {
//         nextQ := [][]int{}
//         for _,p := range q {
//             x,y := p[0],p[1]
//             if grid[x][y] == 0 {
//                 vis[fmt.Sprintf("%d%d",x,y)] = struct{}{}
//                 for _,move := range moves {
//                     r,s := x+move[0] , y+move[1]
//                     _,ok := vis[fmt.Sprintf("%d%d",r,s)] 
//                     if ok {
//                         continue
//                     }
//                     if r < 0 || s < 0 || r > n-1 || s > n-1 {
//                         continue
//                     }
//                     if grid[r][s] == 1 {
//                         continue
//                     }
//                     nextQ = append(nextQ,[]int{r,s})
//                 }
//                 if x == n-1 && y == n-1 {
//                     count++
//                     return count
//                 }
//             }
//         }
//         count++
//         q = nextQ
//     }
//     return -1
// }