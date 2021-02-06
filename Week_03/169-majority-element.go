package Week_03

func majorityElement(nums []int) int {
	return recursion(nums, 0, len(nums)-1)
}

// 多数元素有多种解法，这里使用分治的思想解决
func recursion(nums []int, left, right int) int {
	if left >= right {
		return nums[left]
	}

	mid := left + (right-left)/2

	// 将当前的问题分成两个问题
	l := recursion(nums, left, mid)
	r := recursion(nums, mid+1, right)

	// 如果两边的答案相等，直接返回
	if l == r {
		return l
	}

	// 如果不同，需要统计当前区间下的所有数值，并返回出现次数最多的数字
	lCount := 0
	rCount := 0

	for _, v := range nums[left:right+1] {
		if v == l {
			lCount++
			continue
		}

		if v == r {
			rCount++
		}
	}

	if lCount >= rCount {
		return l
	}

	return r
}
