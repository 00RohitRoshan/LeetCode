func largestOddNumber(num string) string {
    for i := len(num)-1 ; i >= 0; i--{
        // val,_ := strconv.Atoi(string(num[i]))
        if num[i] != '0' && num[i] != '2' && num[i] != '4' && num[i] != '6' && num[i] != '8' {
            return num[:i+1]
        }
    }
    return ""
}