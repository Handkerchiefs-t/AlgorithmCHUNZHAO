package Week_03

func ladderLength(beginWord string, endWord string, wordList []string) int {
	m := map[string]int{}
	g := [][]int{}

	for _, curString := range wordList {
		addEdge(curString, m, &g)
	}

	endId, has := m[endWord]
	if has == false {
		return 0
	}
	beginId := addEdge(beginWord, m, &g)

	maxInt := math.MaxInt64
	dist := make([]int, len(m))
	for i := range dist {
		dist[i] = maxInt
	}
	dist[beginId] = 0

	queue := []int{beginId}
	for len(queue) > 0 {
		curId := queue[0]
		queue = queue[1:]

		if curId == endId {
			return dist[curId]/2 + 1
		}

		for _, edgId := range g[curId] {
			if dist[edgId] == maxInt {
				dist[edgId] = dist[curId] + 1
				queue = append(queue, edgId)
			}
		}
	}

	return 0
}

func addWorld(word string, m map[string]int, g *[][]int) int {
	id, has := m[word]
	if has == false {
		id = len(m)
		m[word] = id
		*g = append(*g, []int{})
	}

	return id
}

func addEdge(word string, m map[string]int, g *[][]int) int {
	reaId := addWorld(word, m, g)

	byteList := []byte(word)

	for i, c := range byteList {
		byteList[i] = '*'
		virId := addWorld(string(byteList), m, g)
		(*g)[virId] = append((*g)[virId], reaId)
		(*g)[reaId] = append((*g)[reaId], virId)
		byteList[i] = c
	}

	return reaId
}

