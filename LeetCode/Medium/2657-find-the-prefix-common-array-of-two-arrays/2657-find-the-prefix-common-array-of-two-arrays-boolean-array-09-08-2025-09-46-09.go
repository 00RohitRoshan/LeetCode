func findThePrefixCommonArray(A []int, B []int) []int {
    n := len(A)
    seen := make([]bool, n+1) // values are 1..n
    ans := make([]int, n)
    count := 0

    for i := 0; i < n; i++ {
        if seen[A[i]] {
            count++
        }
        seen[A[i]] = true

        if seen[B[i]] {
            count++
        }
        seen[B[i]] = true

        ans[i] = count
    }

    return ans
}
