package Week_05

func isValidSudoku(board [][]byte) bool {
	rowList := [9][9]bool{}
	colList := [9][9]bool{}
	bolckList := [9][9]bool{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			n := board[i][j] - '1'
			if rowList[i][n] || colList[j][n] || bolckList[getBolckIndex(i, j)][n] {
				return false
			}

			rowList[i][n] = true
			colList[j][n] = true
			bolckList[getBolckIndex(i, j)][n] = true
		}
	}

	return true
}

func getBolckIndex(i, j int) int {
	return i / 3 * 3 + j / 3
}
