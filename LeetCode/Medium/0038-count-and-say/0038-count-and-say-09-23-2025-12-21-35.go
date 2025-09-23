func countAndSay(n int) string {
    start := "1"
    return find(start,n)
}
func find(s string , n int) string {
    if n == 1 {
        return s
    }
    count := 0
    prev := '0'
    res := []rune{}
    for _,r := range s {
        if r == prev {
            count++
        }else if count != 0 {
            res = append(res,[]rune(strconv.Itoa(count))...)
            res = append(res,prev)
            prev = r
            count = 1
        }else{
            prev = r
            count = 1
        }
    }
    res = append(res,[]rune(strconv.Itoa(count))...)
    res = append(res,prev)
    return find(string(res),n-1)
}