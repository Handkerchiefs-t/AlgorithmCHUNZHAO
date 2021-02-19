package Week_04

func longestValidParentheses(s string) int {
	m := len(s)
	dp := make([]int, m)
	maxLength := 0

	for i := 1; i < m; i++ {
		if s[i] == '(' {
			continue
		}

		if s[i-1] == '(' {
			if i - 2 >= 0 {
				dp[i] = dp[i-2] + 2
			} else {
				dp[i] = 2
			}
		} else {
			if i - dp[i-1] - 1 >= 0 && s[i - dp[i-1] - 1] == '(' {
				if i - dp[i-1] - 2 >= 0 {
					dp[i] = dp[i-1] + dp[i - dp[i-1] - 2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}

		if dp[i] > maxLength {
			maxLength = dp[i]
		}
	}

	return maxLength
}
