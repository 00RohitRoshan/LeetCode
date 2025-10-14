func shortestCommonSupersequence(s string, t string) string {
    n := len(s)
    m := len(t)
    dp := make([][]int,n+1)
    for i := range n+1{
        dp[i] = make([]int,m+1)
    }
    for i,a := range s {
        for j,b := range t {
            if a == b {
                dp[i+1][j+1] = 1+dp[i][j]
            }else{
                dp[i+1][j+1] = max(dp[i][j+1],dp[i+1][j])
            }
        }
    }
    res := []byte{}
    i,j := n,m
    for i > 0 && j > 0 {
        if s[i-1] == t[j-1] {
            res = append(res,s[i-1])
            i--
            j--
        }else if dp[i-1][j] > dp[i][j-1] {
            res = append(res,s[i-1])
            i--
        }else{
            res = append(res,t[j-1])
            j--
        }
    }
    for i > 0 {
        res = append(res,s[i-1])
        i--
    }
    for j > 0 {
        res = append(res,t[j-1])
        j--
    }
    return reverse(res)
}
func reverse(r []byte) string {
    n := len(r)
    for i := range n/2 {
        r[i] , r[n-i-1] = r[n-i-1],r[i]
    }
    return string(r)
}