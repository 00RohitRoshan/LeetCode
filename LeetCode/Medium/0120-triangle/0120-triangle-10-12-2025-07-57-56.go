func minimumTotal(triangle [][]int) int {
    n := len(triangle)
    mem := triangle[n-1]
    for n > 1 {
        for c := range len(triangle[n-2]){
            mem[c] = min(triangle[n-2][c]+mem[c],triangle[n-2][c]+mem[c+1])
        }
        n--
    }
    return mem[0]
}
func find(r int,c int,t [][]int,mem [][]int) int {
    if r == len(t) {
        return 0
    }
    if mem[r][c] != 0 {
        return mem[r][c]
    }
    mem[r][c] = min(t[r][c]+find(r+1,c,t,mem),t[r][c]+find(r+1,c+1,t,mem))
    return mem[r][c]
}