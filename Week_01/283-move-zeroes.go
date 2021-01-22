package Week_01

func moveZeroes(nums []int)  {
	left := 0

	for right, value := range nums {
		if value != 0 {
			nums[right], nums[left] = nums[left], nums[right]
			left++
		}
	}
}
