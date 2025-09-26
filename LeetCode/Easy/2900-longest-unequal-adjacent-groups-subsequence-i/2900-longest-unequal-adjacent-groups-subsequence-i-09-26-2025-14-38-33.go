func getLongestSubsequence(words []string, groups []int) []string {
    last := -1
    res  := []string{}
    for i := range words {
        if groups[i] != last {
            res = append(res,words[i])
            last = groups[i]
        }
    }
    return res
}