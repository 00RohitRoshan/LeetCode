func maximumLength(nums []int, k int) int {
    n := len(nums)
    dp := make([][]int, n+1)
    dp[0] = make([]int, k+1)
    lastIdx := make(map[int]int, n)
    maxSoFar := make([]int, k+1)

    for i := range k+1 {
        maxSoFar[i] = 1
    }

    for i, num := range nums {
        dp[i+1] = make([]int, k+1)
        for j := k; j >= 0; j-- {
            if j == 0 {
                if numDupe, found := lastIdx[num]; found {
                    dp[i+1][j] = dp[numDupe+1][j]+1
                } else {
                    dp[i+1][j]++
                }
            } else if i == 0 {
                dp[i+1][j] = 1
            } else if num == nums[i-1] {
                dp[i+1][j] = max(dp[i][j] + 1, maxSoFar[j-1]+1)
            } else {
                if numDupe, found := lastIdx[num]; found {
                    dp[i+1][j] = dp[numDupe+1][j]
                }

                dp[i+1][j] = max(dp[i+1][j]+1, maxSoFar[j-1]+1)
            }

            maxSoFar[j] = max(maxSoFar[j], dp[i+1][j])
        }
        lastIdx[num] = i
        // fmt.Printf("%2d %v\n",num, dp[i])
    }
    
    return maxSoFar[k]
}