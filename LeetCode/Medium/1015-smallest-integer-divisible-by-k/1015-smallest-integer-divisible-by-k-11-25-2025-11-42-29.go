func smallestRepunitDivByK(k int) int {
    if k & 1 == 0 {
        return -1
    }

    rem := 0
    n := 0
    for length := 1; length <= k; length++ {
        n = n*10 + 1
        // fmt.Println(n)
        rem = (rem*10 + 1) % k
        if rem == 0 {
            return length
        }
    }

    return -1
}