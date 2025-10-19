func generateParenthesis(n int) []string {
    s := make([]byte,2*n)
    res := []string{}
    f(0,0,s,n,&res)
    return res
}
func f(open int,close int,s []byte,n int,res *[]string){
    if open == close && open + close == 2*n {
        *res = append(*res,string(s))
        return
    }
    if open < n {
        s[open+close] = '('
        f(open+1,close,s,n,res)
    }
    if open > close {
        s[open+close] = ')'
        f(open,close+1,s,n,res)
    }
}