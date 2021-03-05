package Week_06

func isMatch(s string, p string) bool {
	s = " " + s
	p = " " + p
	m := len(s)
	n := len(p)

	dp := make([][]bool, m)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	dp[0][0] = true
	for i := 1; i < n; i++ {
		if p[i] == '*' {
			dp[0][i] = dp[0][i-1]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if s[i] == p[j] || p[j] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j] == '*' {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}
