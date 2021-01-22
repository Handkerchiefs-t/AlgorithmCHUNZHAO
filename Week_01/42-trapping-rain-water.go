package Week_01

func trap(height []int) int {
	res := 0

	// 左右指针
	leftPoint := 0
	rightPoint := len(height)-1
	// 左右指针经过的最大高度
	leftMax := 0
	rightMax := 0

	for leftPoint < rightPoint {
		// 如果右边指针较低，代表左边高的柱子可以承接住右边柱子的接的雨水
		if height[leftPoint] > height[rightPoint]{
			if height[rightPoint] < rightMax {  // 如果当前高度比最高的高度低，高度差就是可以接到的雨水
				res += rightMax - height[rightPoint]
			} else {
				rightMax = height[rightPoint]  // 如果当前最高，更新最高值
			}

			rightPoint--
		} else {
			if height[leftPoint] < leftMax {
				res += leftMax - height[leftPoint]
			} else {
				leftMax = height[leftPoint]
			}

			leftPoint++
		}
	}

	return res
}

