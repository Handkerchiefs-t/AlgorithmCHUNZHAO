package Week_06

func lengthOfLIS(nums []int) int {
	m := len(nums)
	dp := make([]int, m)
	for i := range dp {
		dp[i] = 1
	}

	res := 1
	for i := 1; i < m; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		res = max(res, dp[i])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
