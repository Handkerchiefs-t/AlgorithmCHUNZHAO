package Week_03

func updateBoard(board [][]byte, click []int) [][]byte {
	x := click[0]
	y := click[1]

	if board[x][y] == 'M' {
		board[x][y] = 'X'
		return board
	}

	recursion(&board, x, y)

	return board
}

func recursion(board *[][]byte, x, y int) {
	mCount := 0

	for i := range dirX {
		tx := x + dirX[i]
		ty := y + dirY[i]

		if tx < 0 || tx >= len(*board) || ty < 0 || ty >= len((*board)[0]) {
			continue
		}

		if (*board)[tx][ty] == 'M' {
			mCount++
		}
	}

	if mCount > 0 {
		(*board)[x][y] = byte(mCount + '0')
		return
	}

	for i := range dirX {
		(*board)[x][y] = 'B'
		tx := x + dirX[i]
		ty := y + dirY[i]

		if tx < 0 || tx >= len(*board) || ty < 0 || ty >= len((*board)[0]) || (*board)[tx][ty] != 'E' {
			continue
		}

		recursion(board, tx, ty)
	}
}

var dirX = [...]int{0, 1, 0, -1, 1, 1, -1, -1}
var dirY = [...]int{1, 0, -1, 0, 1, -1, 1, -1}

