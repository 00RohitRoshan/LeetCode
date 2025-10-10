func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
    gIdx := 0
    cIdx := 0
    for gIdx < len(g) && cIdx < len(s) {
        if g[gIdx] <= s[cIdx] {
            gIdx++
        }
        cIdx++
    }
    return gIdx
}