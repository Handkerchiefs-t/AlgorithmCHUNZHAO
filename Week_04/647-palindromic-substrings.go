package Week_04

func countSubstrings(s string) int {
	m := len(s)
	count := 0

	dp := make([][]bool, m)
	for i := range dp {
		dp[i] = make([]bool, m)
	}

	for i := 0; i < m; i++ {
		for j := 0; j <= i; j++ {

			if i == j {
				dp[i][j] = true
				count++
			} else if i - 1 == j && s[i] == s[j] {
				dp[i][j] = true
				count++
			} else if s[i] == s[j] && dp[i-1][j+1] {
				dp[i][j] = true
				count++
			}
		}
	}

	return count
}
