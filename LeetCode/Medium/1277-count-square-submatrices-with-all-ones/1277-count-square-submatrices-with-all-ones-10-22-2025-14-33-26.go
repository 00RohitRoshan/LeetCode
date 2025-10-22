func countSquares(matrix [][]int) int {
    m := len(matrix)
    n := len(matrix[0])
    mem := make([][]int,m)
    for i := range m {
        mem[i] = make([]int,n)
        for j := range n {
            mem[i][j] = -1
        }
    }
    ans := 0
    var count func(int,int)int
    count = func(i int,j int) int {
        if mem[i][j] != -1 {
            return mem[i][j]
        }
        if i==0 || j==0 {
            ans += matrix[i][j]
            // fmt.Printf("i:%d j:%d ans:%3d matrix[i][j]:%d\n",i,j,ans,matrix[i][j])
            mem[i][j] = matrix[i][j]
            return mem[i][j]
        }
        il,jl := i,j-1
        it,jt := i-1,j
        id,jd := i-1,j-1
        c := 0
        if matrix[i][j] == 1 {
            c = 1 + min (count(il,jl),count(it,jt),count(id,jd))
            // fmt.Printf("i:%d j:%d ans:%3d c:%d\n",i,j,ans,c)
            ans += c
        }else{
            count(il,jl)
            count(it,jt)
            count(id,jd)
        }
        mem[i][j] = c
        return mem[i][j]
    }
    count(m-1,n-1)
    return ans
}