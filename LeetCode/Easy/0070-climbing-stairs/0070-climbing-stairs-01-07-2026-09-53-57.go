func climbStairs(n int) int {
    prev := 1
    cur := 2
    if n < 3 {
        return n
    }
    for i := 3 ; i<=n ; i++{
        t := prev
        prev = cur
        cur += t
    } 
    return cur
}