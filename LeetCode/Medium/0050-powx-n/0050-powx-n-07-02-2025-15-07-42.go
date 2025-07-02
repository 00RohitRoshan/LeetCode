func myPow(x float64, n int) float64 {
    m := n
    if n < 0{
        x = 1.0/x
        m *= -1
    }
    res := 1.0
    for m>0{
        res *= x
        m--
    }
    return res
}