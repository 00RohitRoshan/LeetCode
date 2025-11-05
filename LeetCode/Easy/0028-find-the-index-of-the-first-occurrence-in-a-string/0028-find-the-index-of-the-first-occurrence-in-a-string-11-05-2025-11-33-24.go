func strStr(haystack string, needle string) int {
    n := len(haystack)
    m := len(needle)
    for i := range n {
        if haystack[i] == needle[0] {
            o := i+m
            if o <= n {
                if haystack[i:o] == needle {
                    return i
                }
            }
        }
    }
    return -1
}