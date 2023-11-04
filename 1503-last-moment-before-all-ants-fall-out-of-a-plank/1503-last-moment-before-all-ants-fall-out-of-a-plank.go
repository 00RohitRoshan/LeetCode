func getLastMoment(n int, left []int, right []int) int {
    maxLeft := 0
    for _, val := range left {
        if val > maxLeft {
            maxLeft = val
        }
    }
    
    minRight := n
    for _, val := range right {
        if val < minRight {
            minRight = val
        }
    }
    
    return int(math.Max(float64(maxLeft), float64(n - minRight)))
}