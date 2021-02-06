package Week_03

func numIslands(grid [][]byte) int {
	landCount := 0

	for x, xList := range grid {
		for y := range xList {
			if grid[x][y] == '1' {
				recursion(&grid, x, y)
				landCount++
			}
		}
	}

	return landCount
}

func recursion(grid *[][]byte, x, y int) {
	if x < 0 || x >= len(*grid) || y < 0 || y >= len((*grid)[0]) {
		return
	}

	if (*grid)[x][y] == '1' {
		(*grid)[x][y] = '0'
	} else {
		return
	}

	recursion(grid, x-1, y)
	recursion(grid, x+1, y)
	recursion(grid, x, y-1)
	recursion(grid, x, y+1)
}

