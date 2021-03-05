package Week_06

func longestPalindrome(s string) string {
	m := len(s)
	dp := make([][]int, m)
	length := 0
	res := ""

	for i := range dp {
		dp[i] = make([]int, m)
	}

	for i := 0; i < m; i++ {
		for j := 0; j <= i; j++ {
			if i == j {
				dp[i][j] = 1
			} else if i - j == 1 && s[i] == s[j] {
				dp[i][j] = 2
			} else if s[i] == s[j] && dp[i-1][j+1] != 0 {
				dp[i][j] = dp[i-1][j+1] + 2
			}

			if dp[i][j] > length {
				length = i-j+1
				res = s[j:i+1]
			}
		}
	}

	return res
}
