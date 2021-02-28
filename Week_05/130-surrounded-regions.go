package Week_05

func solve(board [][]byte)  {
	if len(board) == 0 {
		return
	}

	m := len(board)
	n := len(board[0])
	virtualPoint := m * n
	parent := make([]int, virtualPoint+1)
	for i := range parent {
		parent[i] = i
	}

	_find := func (i int) int {
		root := i

		for i != parent[i] {
			i = parent[i]
		}

		for root != parent[root] {
			tmp := root
			root = parent[root]
			parent[tmp] = i
		}

		return i
	}

	_union := func (i, j int) {
		p1 := _find(i)
		p2 := _find(j)

		parent[p1] = p2
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				if i <= 0 || i >= m-1 || j <= 0 || j >= n-1 {
					_union( i * n + j, virtualPoint)
				} else {
					for f := range fontX {
						tryX := fontX[f] + i
						tryY := fontY[f] + j

						if board[tryX][tryY] == 'O' {
							_union(i * n + j, tryX * n + tryY)
						}
					}
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && _find(i*n+j) != _find(virtualPoint) {
				board[i][j] = 'X'
			}
		}
	}

}

var fontX = [...]int{1, -1, 0, 0}
var fontY = [...]int{0, 0, 1, -1}
