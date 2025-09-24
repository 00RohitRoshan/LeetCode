func fractionToDecimal(numerator int, denominator int) string {
    if numerator == 0 {
        return "0"
    }
    n := int64(abs(numerator)) // for handling overflow
    d := int64(abs(denominator))
    var sb  strings.Builder
    if (numerator<0) != (denominator<0) {
        sb.WriteString("-")
    }
    sb.WriteString(fmt.Sprintf("%d",n/d)) // also works with it instead of strconv.FormatInt(i,10) 
    rem := n%d
    if rem == 0 {
        return sb.String()
    }
    sb.WriteString(".")
    has := map[int64]int{rem:sb.Len()}
    for rem != 0 {
        rem *= 10    // simulate long division important
        pos , ok := has[rem]
        if ok {
            s := sb.String()
            return s[:pos] + "(" + s[pos:] + ")"
        }
        has[rem] = sb.Len()
        sb.WriteString(fmt.Sprintf("%d",rem/d))
        rem = rem%d
    }
    return sb.String()
}

func abs(i int)int{
    if i < 0 {
        return -i
    }
    return i
}