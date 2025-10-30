func digArtifacts(n int, artifacts [][]int, dig [][]int) int {
    digged := make([][]int,n)
    for i := range n {
        digged[i] = make([]int,n)
    }
    for _,d := range dig {
        digged[d[0]][d[1]] = 1
    }
    count := 0
    for _,a := range artifacts {
        for i := a[0]; i <= a[2]; i++ {
            for j := a[1]; j <= a[3]; j++ {
                if digged[i][j] == 0 {
                    goto skip
                }
                if i == a[2] && j == a[3] {
                    count++
                }
            }
        }
        skip:
    }
    return count
}