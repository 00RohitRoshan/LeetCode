func solve(board [][]byte)  {
    m := len(board)
    n := len(board[0])
    for i := range m{
        if board[i][0] == 'O' {
            dfs(i,0,board)
        }
        if board[i][n-1] == 'O' {
            dfs(i,n-1,board)
        }
    }
    for i := range n{
        if board[0][i] == 'O' {
            dfs(0,i,board)
        }
        if board[m-1][i] == 'O' {
            dfs(m-1,i,board)
        }
    }
    for i := range m {
        for j := range n {
            if board[i][j] == 'O' {
                board[i][j] = 'X'
            }
            if board[i][j] == 'r' {
                board[i][j] = 'O'
            }
        }
    }
}
func dfs(row int, col int, board [][]byte){
    m := len(board)
    n := len(board[0])
    fmt.Println("inside dfs",row,col)
    if board[row][col] == 'O' {
        board[row][col] = 'r'
        if row < m-1 {
            dfs(row+1,col,board)
        }
        if row > 0 {
            dfs(row-1,col,board)
        }
        if col < n-1 {
            dfs(row,col+1,board)
        }
        if col > 0 {
            dfs(row,col-1,board)
        }
    }
}