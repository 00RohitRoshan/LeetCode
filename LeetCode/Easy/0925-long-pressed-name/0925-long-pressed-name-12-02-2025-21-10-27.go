func isLongPressedName(name string, typed string) bool {
    n := 0
    t := 0
    for t < len(typed) {
        if n < len(name) && name[n] == typed[t] {
            n++
            t++
        }else if t > 0 && typed[t] == typed[t-1] {
            t++
        }else{
            return false
        }
    }
    return n == len(name)
}