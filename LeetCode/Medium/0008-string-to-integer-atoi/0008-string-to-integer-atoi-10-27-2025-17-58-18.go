func myAtoi(s string) int {
    i := 0
    n := len(s)
    for i < n && s[i] == ' '{
        i++
    }
    sign := 1
    if i < n &&  s[i] == '-' {
        sign = -1
        i++
    }else if  i < n && s[i] == '+' {
            i++
        }
    num := 0
    for i < n && s[i] >= '0' && s[i] <= '9'{
        num = (num*10)+int(s[i]-'0')
        if num > math.MaxInt32 {
            if sign == -1 {
                return (sign*math.MaxInt32)-1
            } 
            return sign*math.MaxInt32
        }
        i++
    }
    return num*sign
}