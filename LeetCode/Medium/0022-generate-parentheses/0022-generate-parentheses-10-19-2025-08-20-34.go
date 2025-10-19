func generateParenthesis(n int) []string {
    s := []byte{}
    res := []string{}
    f(0,0,&s,n,&res)
    return res
}
func f(open int,close int,s *[]byte,n int,res *[]string){
    if open == close && open + close == 2*n {
        *res = append(*res,string(*s))
        return
    }
    if open < n {
        o := append(*s,'(')
        f(open+1,close,&o,n,res)
    }
    if close < n && open > close {
        c := append(*s,')')
        f(open,close+1,&c,n,res)
    }
}