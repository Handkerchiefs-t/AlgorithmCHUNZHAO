package Week_03

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	// 初始化 string -> id 的哈希表
	strToId := map[string]int{}
	for i, s := range wordList {
		strToId[s] = i
	}
	if _, ok := strToId[beginWord]; ok == false {
		wordList = append(wordList, beginWord)
		strToId[beginWord] = len(wordList) - 1
	}
	if _, ok := strToId[endWord]; ok == false {
		return [][]string{}
	}

	// 建图，edgs[i] 可以获得 i 可以到达的节点，也就是它可以接龙的字符
	edgs := make([][]int, len(wordList))
	for i := 0; i < len(wordList); i++ {
		for j := i; j < len(wordList); j++ {
			if transform(wordList[i], wordList[j]) == true {
				edgs[i] = append(edgs[i], j)
				edgs[j] = append(edgs[j], i)
			}
		}
	}

	// 定义 返回列表、路径花费表、队列
	res := [][]string{}
	cost := make([]int, len(wordList))
	queue := [][]int{[]int{strToId[beginWord]}}

	// 初始化路径花费表
	x := math.MaxInt64
	for i := range cost {
		cost[i] = x
	}
	cost[strToId[beginWord]] = 0

	// 广度优先搜索
	for i:=0; i < len(queue); i++ {
		pathList := queue[i]
		lastId := pathList[len(pathList)-1]

		// 找到了结束点，添加路径到 res
		if wordList[lastId] == endWord {
			tmp := make([]string, 0, len(pathList))

			for _, s := range pathList {
				tmp = append(tmp, wordList[s])
			}

			res = append(res, tmp)
		} else {
			// 不是结束点，向队列中放路径
			for _, nextEdgs := range edgs[lastId] {
				// 到达 next 的代价变小（或不变），将这个路径加入进去。如果新的路径的代价大于之前的代价，无视即可
				if cost[lastId] + 1 <= cost[nextEdgs] {
					cost[nextEdgs] = cost[lastId]+1
					tmp := append([]int{}, pathList...)
					tmp = append(tmp, nextEdgs)
					queue = append(queue, tmp)
				}
			}
		}
	}

	return res
}

// 判断 from 能否转化为 to
func transform(from ,to string) bool {
	for i := range from {
		if from[i] != to[i] {
			return from[i+1:] == to[i+1:]
		}
	}

	return false
}
