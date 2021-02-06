package Week_03

func canJump(nums []int) bool {
	maxDistance := 0

	for i, v := range nums {
		if i > maxDistance {
			return false
		}

		if i + v > maxDistance {
			maxDistance = i + v
		}
	}

	return true
}

