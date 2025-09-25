type Trip struct {
    Row, Col int
}

func shortestPathBinaryMatrix(grid [][]int) int {
    if grid[0][0] != 0 || grid[len(grid)-1][len(grid[0])-1] != 0 {
        return -1
    }
    
    queue := []Trip{{0,0}}
    grid[0][0] = 1
        
    for len(queue) > 0 {
        now := queue[0]
        queue = queue[1:]
        
        if now.Row == len(grid)-1 && now.Col == len(grid[0])-1 {
            return grid[now.Row][now.Col]
        }

        moves:= [][]int{{-1,-1},{-1,0},{-1,1},{0,-1},{0,1},{1,-1},{1,0},{1,1}}
        
        for _, move := range moves {
            t := Trip{now.Row+move[0], now.Col+move[1]}
            
            if t.Row < 0 || t.Row >= len(grid) || t.Col < 0 || t.Col >= len(grid[0]) || grid[t.Row][t.Col] != 0 {
                continue
            }
            
            grid[t.Row][t.Col] = grid[now.Row][now.Col]+1
            queue = append(queue, t)
        }
    }
    
    return -1
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