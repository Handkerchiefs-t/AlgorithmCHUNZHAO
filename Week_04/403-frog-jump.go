package Week_04

func canCross(stones []int) bool {
	m := len(stones)

	// 路不要走窄了，谁说dp必须用数组的？
	hash := map[int]map[int]bool{}
	for _, v := range stones {
		hash[v] = map[int]bool{}
	}

	hash[0][0] = true

	for _, stone := range stones {
		for step, _ := range hash[stone] {
			for i := step - 1; i <= step + 1; i++ {
				if stepSet, has := hash[stone+i]; has {
					stepSet[i] = true
				}
			}
		}
	}

	if len(hash[stones[m-1]]) > 0 {
		return true
	}
	return false
}
