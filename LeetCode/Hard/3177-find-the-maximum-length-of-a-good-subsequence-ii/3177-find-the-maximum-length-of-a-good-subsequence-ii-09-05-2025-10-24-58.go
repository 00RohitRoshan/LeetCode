func maximumLength(nums []int, k int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// prevIdx[i] = last index of nums[i] that occurs BEFORE i, or -1 if none.
	prevIdx := make([]int, n)
	last := make(map[int]int, n)
	for i := 0; i < n; i++ {
		if p, ok := last[nums[i]]; ok {
			prevIdx[i] = p
		} else {
			prevIdx[i] = -1
		}
		last[nums[i]] = i
	}

	type key struct{ i, j int }
	memoF := make(map[key]int) // f(i,j)
	memoG := make(map[key]int) // g(i,j)

	var f func(i, j int) int
	var g func(i, j int) int

	// f(i,j): equivalent to dp[i+1][j] in your bottom-up version:
	// best length ending exactly at index i with up to j "ops".
	f = func(i, j int) int {
		if i < 0 {
			return 0
		}
		if j < 0 {
			return 0
		}
		K := key{i, j}
		if v, ok := memoF[K]; ok {
			return v
		}

		var res int
		if j == 0 {
			// Same as:
			// if num seen before -> dp[prev+1][0] + 1
			// else -> 1
			p := prevIdx[i]
			if p != -1 {
				res = f(p, 0) + 1
			} else {
				res = 1
			}
		} else if i == 0 {
			// else-if i == 0 -> 1
			res = 1
		} else if nums[i] == nums[i-1] {
			// num == nums[i-1] -> max(dp[i][j] + 1, maxSoFar[j-1] + 1)
			res = max(f(i-1, j)+1, g(i-1, j-1)+1)
		} else {
			// else:
			// if seen before: base = dp[prev+1][j], else base = 0
			// dp = max(base + 1, maxSoFar[j-1] + 1)
			base := 0
			if p := prevIdx[i]; p != -1 {
				base = f(p, j)
			}
			res = max(base+1, g(i-1, j-1)+1)
		}

		memoF[K] = res
		return res
	}

	// g(i,j): equivalent to maxSoFar[j] after processing up to index i.
	// g(i,j) = max(g(i-1,j), f(i,j))
	g = func(i, j int) int {
		if i < 0 {
			return 0
		}
		K := key{i, j}
		if v, ok := memoG[K]; ok {
			return v
		}
		res := max(g(i-1, j), f(i, j))
		memoG[K] = res
		return res
	}

	return g(n-1, k)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}