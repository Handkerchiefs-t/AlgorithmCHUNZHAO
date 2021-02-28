package Week_05

func solveSudoku(board [][]byte)  {
	rowList := [9][9]bool{}
	colList := [9][9]bool{}
	bolckList := [9][9]bool{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			n := board[i][j] - '1'

			rowList[i][n] = true
			colList[j][n] = true
			bolckList[getBolckIndex(i, j)][n] = true
		}
	}

	recursion(board, 0, 0, rowList, colList, bolckList)
}

func recursion(board [][]byte, i, j int, rowList, colList, bolckList [9][9]bool) bool {
	for ; i < 9; i++ {
		for ; j < 9; j++ {
			if board[i][j] != '.' {
				continue
			}

			for char := byte('1'); char <= '9'; char++ {
				num := int(char - '1')

				if rowList[i][num] || colList[j][num] || bolckList[getBolckIndex(i, j)][num] {
					continue
				}

				rowList[i][num] = true
				colList[j][num] = true
				bolckList[getBolckIndex(i, j)][num] = true
				board[i][j] = char

				if recursion(board, i, j, rowList, colList, bolckList) {
					return true
				}

				rowList[i][num] = false
				colList[j][num] = false
				bolckList[getBolckIndex(i, j)][num] = false
				board[i][j] = '.'
			}

			return false
		}
		j = 0
	}

	return true
}

func getBolckIndex(i, j int) int {
	return i / 3 * 3 + j / 3
}
