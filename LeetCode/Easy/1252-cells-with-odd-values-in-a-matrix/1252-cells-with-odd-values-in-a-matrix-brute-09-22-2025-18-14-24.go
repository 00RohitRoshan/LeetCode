func oddCells(m int, n int, indices [][]int) int {
    mat := make([][]int,m)
    for i := range m {
        mat[i] = make([]int,n)
    } 
    for _,v := range indices {
        r := v[0]
        c := v[1]
        for i := range m {
            if i == r {
                for j := range n {
                    mat[i][j] += 1
                }
            }
            mat[i][c] += 1
        }
    }
    count := 0
    for i := range m {
        for j := range n {
            if mat[i][j]%2 != 0 {
                count++
            }
        }
    }
    return count
}