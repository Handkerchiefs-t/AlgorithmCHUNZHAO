package Week_05

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	parent := make([]int, n)
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

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isConnected[i][j] == 1 {
				_union(i, j)
			}
		}
	}

	// fmt.Println(parent)

	count := 0
	for i, v := range parent {
		if i == v {
			count++
		}
	}

	return count
}
