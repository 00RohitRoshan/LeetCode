func solveNQueens(n int) [][]string {
    ans := [][]string{}
    board := make([][]rune,n)
    for i := range n {
        board[i] = make([]rune,n)
        for j := range n {
            board[i][j] = '.'
        }
    }
    left := make([]int,2*n)
    topD := make([]int,2*n)
    bottomD := make([]int,2*n)
    solve(0,left,topD,bottomD,board,&ans)
    return ans
}
func solve(col int,left []int,top []int,bottom []int,board [][]rune,ans *[][]string){
    n := len(board)
    if col == n {
        *ans = append(*ans,runeToString(board))
        return
    }
    for i := range n {
        if valid(col,i,n,left,top,bottom) {
            mark(col,i,n,left,top,bottom,board)
            solve(col+1,left,top,bottom,board,ans)
            clear(col,i,n,left,top,bottom,board)
        }
    }
}
func valid(col int,row int,n int,left []int,top []int,bottom []int)bool{
    return left[row] == 0 && top[col+row] == 0 && bottom[(n-1)+(row-col)] == 0
}
func mark(col int,row int,n int,left []int,top []int,bottom []int,board [][]rune){
    left[row] = 1 
    top[col+row] = 1 
    bottom[(n-1)+(row-col)] = 1
    board[row][col] = 'Q'
}
func clear(col int,row int,n int,left []int,top []int,bottom []int,board [][]rune){
    left[row] = 0 
    top[col+row] = 0 
    bottom[(n-1)+(row-col)] = 0
    board[row][col] = '.'
}
func runeToString(board [][]rune) []string {
    n := len(board)
    res := make([]string,n)
    for i := range n {
        res[i] = string(board[i])
    }
    return res
}