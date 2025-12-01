func minNumberOperations(target []int) int {
    n := len(target)
    sum := 0
    start := 0 
    i := 0
    for i < n {
        for i<n && target[i] >= start {
            start = target[i]
            i++
        }
        // fmt.Println(start)
        sum += start
        for i<n && target[i] <= start {
            start = target[i]
            i++
        }
        if i != n {
            // fmt.Println(n,i,-start)
            sum -= start
        }
    }
    return sum
}