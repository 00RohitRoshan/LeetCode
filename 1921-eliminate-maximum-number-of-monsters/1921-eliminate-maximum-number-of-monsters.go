func eliminateMaximum(dist []int, speed []int) int {
    n := len(dist)
    time_to_city := make([]float64, n)
    
    for i := 0; i < n; i++ {
        time_to_city[i] = float64(dist[i]) / float64(speed[i])
    }
    
    sort.Float64s(time_to_city)
    
    for i := 0; i < n; i++ {
        if time_to_city[i] <= float64(i) {
            return i
        }
    }
    
    return n
}