func diagonalSum(mat [][]int) int {
    n := len(mat)
    sum := 0
    for i := 0; i < n; i++ {
        sum += mat[i][i] // primary diagonal
        if i != n-i-1 {  // avoid double-counting the center when n is odd
            sum += mat[i][n-i-1]
        }
    }
    return sum
}
