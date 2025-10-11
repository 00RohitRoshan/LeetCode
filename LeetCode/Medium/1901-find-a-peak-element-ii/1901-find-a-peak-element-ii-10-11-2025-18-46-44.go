func findPeakGrid(mat [][]int) []int {
    m := len(mat)
    n := len(mat[0])
    maxi := []int{0,0}
    for i := range m {
        for j := range n {
            if mat[i][j] > mat[maxi[0]][maxi[1]] {
                maxi[0] = i
                maxi[1] = j
            }
        }
    }
    return maxi
}