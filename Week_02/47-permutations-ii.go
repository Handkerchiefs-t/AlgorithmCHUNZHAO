package Week_02

// 方法1：使用集合去重
// 方法2：排序+跳跃去重，这里用这个方法

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)  // 排序，后续去重需要
	res := [][]int{}
	used := make([]bool, len(nums))

	recursion(nums, []int{}, used, &res)

	return res
}

func recursion(nums, path []int, used []bool, res *[][]int) {
	if len(nums) == len(path) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	for i := range nums {
		// 出现重复的情况，直接跳过
		if i > 0 && used[i-1] == false && nums[i] == nums[i-1] {
			continue
		}

		if used[i] == false {
			used[i] = true
			path = append(path, nums[i])
			recursion(nums, path, used, res)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
}
