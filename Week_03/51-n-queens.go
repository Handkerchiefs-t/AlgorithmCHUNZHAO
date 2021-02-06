package Week_03

var x = struct{}{}

func solveNQueens(n int) [][]string {
	resIntList := [][]int{}

	recursion(n, 0, []int{}, map[int]struct{}{}, map[int]struct{}{}, map[int]struct{}{}, &resIntList)

	return translationToString(resIntList)
}

func recursion(n, row int, path []int, colSet, pieSet, naSet map[int]struct{}, res *[][]int) {
	if n == row {
		*res = append(*res, append([]int{}, path...))
		return
	}

	for col := 0; col < n; col++ {
		_, ok1 := colSet[col]
		_, ok2 := pieSet[row + col]
		_, ok3 := naSet[row - col]

		if ok1 || ok2 || ok3 {
			continue
		}

		colSet[col] = x
		pieSet[row + col] = x
		naSet[row - col] = x
		path = append(path, col)

		recursion(n, row+1, path, colSet, pieSet, naSet, res)

		path = path[:len(path)-1]
		delete(colSet, col)
		delete(pieSet, row+col)
		delete(naSet, row-col)
	}
}

func translationToString(in [][]int) [][]string {
	res := [][]string{}

	for _, path := range in {
		queenString := []string{}
		for _, item := range path {
			oneQueen := ""
			for i := 0; i < len(path); i++ {
				if i == item {
					oneQueen = oneQueen + "Q"
					continue
				}
				oneQueen = oneQueen + "."
			}
			queenString = append(queenString, oneQueen)
		}
		res = append(res, queenString)
	}

	return res
}

