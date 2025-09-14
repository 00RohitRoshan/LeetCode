func numEnclaves(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    for i := range m {
        if grid[i][0] == 1 {
            dfs(i,0,grid)
        }
        if grid[i][n-1] == 1 {
            dfs(i,n-1,grid)
        }
    }
    for i := range n {
        if grid[0][i] == 1 {
            dfs(0,i,grid)
        }
        if grid[m-1][i] == 1 {
            dfs(m-1,i,grid)
        }
    }
    count := 0
    for i := range m {
        for j := range n {
            if grid[i][j] == 1 {
                count++
            }
        }
    }
    return count
}

func dfs(row int, col int, grid [][]int){
    m := len(grid)
    n := len(grid[0])
    if grid[row][col] == 1{
        grid[row][col] = 0 
        if row > 0{
            dfs(row-1,col,grid)
        }
        if row < m-1 {
            dfs(row+1,col,grid)
        }
        if col > 0{
            dfs(row,col-1,grid)
        }
        if col < n-1 {
            dfs(row,col+1,grid)
        }
    }
}