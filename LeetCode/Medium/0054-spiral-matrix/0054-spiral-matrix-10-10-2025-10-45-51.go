func spiralOrder(matrix [][]int) []int {
    m := len(matrix)
    if m == 0 {
        return []int{}
    }
    if m == 1 {
        return matrix[0]
    }
    n := len(matrix[0])
    res := make([]int,m*n)
    pos := 0
    top,bottom,left,right := 0,m-1,0,n-1
    //  fmt.Printf("top %d bottom %d left %d right %d \n",top,bottom,left,right)
    for bottom >= top || right >= left {
        for i := left; i<=right; i++{
            if top <= bottom {
                res[pos] = matrix[top][i]
                //  fmt.Printf("%d ",res[pos])
                pos++
            }
        }
        top++
        //  fmt.Printf("\ntop %d \n",top)
        for i := top; i<=bottom; i++ {
            if right >= left {
                res[pos] = matrix[i][right]
                //  fmt.Printf("%d ",res[pos])
                pos++
            }
        }
        right--
        //  fmt.Printf("\n right %d \n",right)
        for i := right; i>=left; i--{
            if bottom >= top {
                res[pos] = matrix[bottom][i]
                //  fmt.Printf("%d ",res[pos])
                pos++
            }
        }
        bottom--
        //  fmt.Printf("\n bottom %d \n",bottom)
        for i := bottom; i>=top; i-- {
            if left <= right {
                res[pos] = matrix[i][left]
                //  fmt.Printf("%d ",res[pos])
                pos++
            }
        }
        left++
        //  fmt.Printf("\n left %d \n",left)
    }
    return res
}