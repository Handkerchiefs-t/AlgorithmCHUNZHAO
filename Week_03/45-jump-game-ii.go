package Week_03

func jump(nums []int) int {
	maxDistance := 0
	curMaxDis := 0
	step := 0

	for i := 0; i < len(nums)-1; i++ {
		if i + nums[i] > maxDistance {
			maxDistance = i + nums[i]
		}

		if curMaxDis == i {
			curMaxDis = maxDistance
			step++
		}
	}

	return step
}

