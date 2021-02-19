package Week_04

func minDistance(word1 string, word2 string) int {
	word1 = " " + word1
	word2 = " " + word2
	m := len(word1)
	n := len(word2)

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		dp[i][0] = i
	}

	for i := 0; i < n; i++ {
		dp[0][i] = i
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if word1[i] == word2[j] {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1])
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}

	return dp[m-1][n-1]
}

func min(arr ...int) int {
	r := arr[0]

	for _, v := range arr {
		if v < r {
			r = v
		}
	}

	return r
}

