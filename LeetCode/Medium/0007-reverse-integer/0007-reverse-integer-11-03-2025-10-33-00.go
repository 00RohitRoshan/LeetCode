func reverse(x int) int {
    sign := 1 
    if x < 0 {
        sign = -1
        x *= -1
    }
    res := 0
    for x > 0 {
        m := x%10
        res = (res*10)+m
        if res >= math.MaxInt32 || res*sign <= math.MinInt32 {
            return 0
        }
        x = (x-m)/10
    }
    return res*sign
}