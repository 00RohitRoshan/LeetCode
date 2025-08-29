func removeDuplicates(s string) string {
    // sRune := rune()
    stack := []byte{}
    n := len(s)
    for i := range n {
        nstack := len(stack)
        var cur  byte
        if nstack > 0 {
            cur = stack[nstack-1]
        }
        // fmt.Println(nstack)
        // fmt.Println("i",i)
        if cur == s[i] {
            stack = stack[:nstack-1]
        } else {
            stack = append(stack,s[i])
        }
    }
    return string(stack)
}