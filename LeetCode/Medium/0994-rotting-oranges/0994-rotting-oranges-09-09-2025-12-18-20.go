func orangesRotting(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    fresh := 0
    queue := [][]int{}

    // Count fresh oranges and collect rotten ones
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                fresh++
            }
            if grid[i][j] == 2 {
                queue = append(queue, []int{i, j})
            }
        }
    }

    if fresh == 0 {
        return 0
    }

    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    minutes := -1

    // BFS
    for len(queue) > 0 {
        next := [][]int{}
        for _, point := range queue {
            x, y := point[0], point[1]
            for _, d := range dirs {
                nx, ny := x+d[0], y+d[1]
                if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == 1 {
                    grid[nx][ny] = 2
                    fresh--
                    next = append(next, []int{nx, ny})
                }
            }
        }
        queue = next
        minutes++
    }

    if fresh > 0 {
        return -1
    }
    return minutes
}
