func maxProduct(nums []int) int64 {
	// find maximum number in nums
	maxN := 0
	for _, x := range nums {
		if x > maxN {
			maxN = x
		}
	}

	// most significant bit
	msb := int(math.Log2(float64(maxN)))
	maxMask := (1 << (msb + 1)) - 1

	// dp array
	dp := make([]int, maxMask+1)
	for _, x := range nums {
		dp[x] = x
	}

	// SOS DP (subset optimization)
	for b := 0; b <= msb; b++ {
		for mask := 0; mask < maxMask; mask++ {
			if mask&(1<<b) != 0 {
				if dp[mask^ (1<<b)] > dp[mask] {
					dp[mask] = dp[mask^(1<<b)]
				}
			}
		}
	}

	var ans int64 = 0
	for _, n := range nums {
		prod := int64(n) * int64(dp[maxMask^n])
		if prod > ans {
			ans = prod
		}
	}
	return ans
}