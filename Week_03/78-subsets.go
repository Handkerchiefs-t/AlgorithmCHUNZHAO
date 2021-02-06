package Week_03

func subsets(nums []int) [][]int {
	res := [][]int{}

	recursion(nums, []int{}, 0, &res)

	return res
}

func recursion(nums, path []int, cur int, res *[][]int) {
	if len(nums) == cur {
		*res = append(*res, append([]int{}, path...))
		return
	}

	recursion(nums, path, cur+1, res)

	path = append(path, nums[cur])
	recursion(nums, path, cur+1, res)
}

