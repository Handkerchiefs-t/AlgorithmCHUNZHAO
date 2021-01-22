package Week_01

func removeDuplicates(nums []int) int {
	slow := 0  // 慢指针，这个指针之前的数都是不重复的
	lastValue := -100  // 记录上一个值，-100是随便选的，这里其实不严谨

	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != lastValue {
			lastValue = nums[fast]
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
		}
	}

	return slow
}
