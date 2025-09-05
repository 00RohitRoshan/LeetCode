func diagonalSum(mat [][]int) int {
    n := len(mat)
    sum := 0
    for i := range n{
        sum += mat[i][i]
    }
    for i := range n {
        if i != n-i-1 {
            sum += mat[i][n-i-1]
        }
    }
    return sum
}