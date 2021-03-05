package Week_06

func racecar(target int) int {
	if target == 1 {
		return 1
	}

	dp := make([]int, target+1)
	dp[0] = 0
	dp[1] = 1

	noBackStep := 2
	maxDistance := 3

	for i := 2; i <= target; i++ {
		if i == maxDistance {
			dp[maxDistance] = noBackStep
			noBackStep++
			maxDistance = ( 1 << noBackStep) - 1

			continue
		}

		dp[i] = noBackStep + 1 + dp[maxDistance-i]

		for back := 0; back < noBackStep-1; back++ {
			needDistance := i + (1 << back) - (1 << (noBackStep - 1))
			tmpStep := noBackStep - 1 + 2 + back + dp[needDistance]
			dp[i] = min(dp[i], tmpStep)
		}
	}

	return dp[target]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
