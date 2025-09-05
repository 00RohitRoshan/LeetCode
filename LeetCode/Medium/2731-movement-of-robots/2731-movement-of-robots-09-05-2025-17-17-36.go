func sumDistance(nums []int, s string, d int) int {
    n := len(s)
    for i, c := range s {
        if c == 'R' {
            nums[i] += d
        } else {
            nums[i] -= d
        }
    }

    sort.Ints(nums)
    prefix := make([]int64, n+1)
    for i := 0; i < n; i++ {
        prefix[i+1] = prefix[i] + int64(nums[i])
    }

    var total int64
    const mod int64 = 1_000_000_007
    for i := 0; i < n; i++ {
        rightSum := prefix[n] - prefix[i+1]
        count := int64(n - i - 1)
        total += rightSum - count*int64(nums[i])
        total %= mod
    }

    return int(total)
}
