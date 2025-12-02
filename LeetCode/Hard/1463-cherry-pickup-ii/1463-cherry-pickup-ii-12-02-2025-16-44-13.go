func cherryPickup(grid [][]int) int {
    row := len(grid)
    col := len(grid[0])
    dp := make([][][]int,row)
    for i := range row {
        dp[i] = make([][]int,col) 
        for j:= range col {
            dp[i][j] = make([]int,col)
            for k := range col {
                dp[i][j][k] = -1
            }
        }
    }
    // fmt.Println(row,col,dp)
    return Pick(0,0,col-1,grid,dp,row-1,col-1)
}
func Pick(r int,c1 int,c2 int,grid [][]int,dp [][][]int,row int,col int)int{
    if r > row || c1 < 0 || c2 < 0 || c1 > col || c2 > col {
        return 0
    }
    if r == row && c1 == c2 {
        return grid[r][c1]
    }
    if dp[r][c1][c2] != -1 {
        return dp[r][c1][c2]
    }
    cur := grid[r][c1]+grid[r][c2]
    if c1 == c2 {
        cur = grid[r][c1]
    }
    maxCherry := cur
    for i := -1; i<2; i++{
        for j := -1; j<2; j++{
            cherry := cur + Pick(r+1,c1+i,c2+j,grid,dp,row,col)
            maxCherry = max(maxCherry,cherry)
        }
    }
    dp[r][c1][c2] = maxCherry
    return maxCherry
}