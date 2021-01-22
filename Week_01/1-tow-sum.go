package Week_01

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	r := make([]int, 0, 2)

	for index, item := range nums {
		if v, ok := m[target - item]; ok {
			r = append(r, index)
			r = append(r, v)
			return r
		} else {
			m[item] = index
		}
	}

	return r
}
