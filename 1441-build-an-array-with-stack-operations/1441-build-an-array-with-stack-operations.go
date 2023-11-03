func buildArray(target []int, n int) []string {
    targetSet := make(map[int]bool)
    for _, num := range target {
        targetSet[num] = true
    }

    var result []string
    for i := 1; i <= target[len(target)-1]; i++ {
        if targetSet[i] {
            result = append(result, "Push")
        } else {
            result = append(result, "Push")
            result = append(result, "Pop")
        }
    }
    return result
}