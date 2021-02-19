package Week_04

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}

	m := len(s)
	dp := make([]int, m+1)

	dp[0] = 1
	dp[1] = 1

	for i := 1; i < m; i++ {
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				dp[i+1] = dp[i-1]
			} else {
				return 0
			}
		} else {
			if s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6') {
				dp[i+1] = dp[i] + dp[i-1]
			} else {
				dp[i+1] = dp[i]
			}
		}
	}

	return dp[m]
}
