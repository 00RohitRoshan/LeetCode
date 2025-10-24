func rotate(matrix [][]int)  {
    m := len(matrix)
    n := len(matrix[0])
    for i := range m {
        for j := range i {
            matrix[i][j],matrix[j][i] = matrix[j][i],matrix[i][j]
        }
    }
    for i := range m {
        for j := range n/2 {
            matrix[i][j],matrix[i][n-j-1] = matrix[i][n-j-1],matrix[i][j]
        }
    }
}