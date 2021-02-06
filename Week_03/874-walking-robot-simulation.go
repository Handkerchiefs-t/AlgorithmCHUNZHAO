package Week_03

func robotSim(commands []int, obstacles [][]int) int {
	// 定义并建立一个包含所有障碍物的集合
	m := map[[2]int]struct{}{}
	for _, v := range obstacles {
		m[[2]int{v[0], v[1]}] = x
	}

	// 定义 当前位置、机器人朝向、最大距离
	nowX := 0
	nowY := 0
	nowFoward := 0
	maxDistance := 0

	for _, v := range commands {
		// 转向
		if v == -1 {
			nowFoward = (nowFoward + 1) % 4
			continue
		}
		if v == -2 {
			nowFoward = (nowFoward + 3) % 4
			continue
		}

		// 前进
		for i := 0; i < v; i++ {
			// 尝试前进到 tx, ty 点
			tx := nowX + fowardX[nowFoward]
			ty := nowY + fowardY[nowFoward]

			// 如果这个点是障碍物，break
			if _, ok := m[[2]int{tx, ty}]; ok {
				break
			}

			// 更新位置和最大距离
			nowX = tx
			nowY = ty
			if tx*tx + ty*ty > maxDistance {
				maxDistance = tx*tx + ty*ty
			}
		}
	}

	return maxDistance
}

var x = struct{}{}
var fowardX = [4]int{0, 1, 0, -1}
var fowardY = [4]int{1, 0, -1, 0}

