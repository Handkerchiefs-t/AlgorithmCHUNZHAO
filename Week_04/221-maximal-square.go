package Week_04

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])
	res := 0

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		dp[i][0] = int(matrix[i][0] - '0')
		res = max(res, dp[i][0])
	}

	for i := 0; i < n; i++ {
		dp[0][i] = int(matrix[0][i] - '0')
		res = max(res, dp[0][i])
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '0' {
				continue
			}

			dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			res = max(res, dp[i][j])
		}
	}

	// fmt.Println(dp)

	return res * res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
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
