func countTriples(n int) int {
    has := make(map[int]bool)
    count := 0
    for i := range n+1 {
        t := i*i
        for j := range i {
            h := j*j
            if has[t-h] {
                if t - h != h {
                    count ++
                }else{
                    count++
                }
            }
        }
        has[t] = true
    }
    return count
}