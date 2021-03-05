package Week_06

func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, s, e int) int {
	if s >= e {
		return 0
	}

	m := (s + e) >> 1
	n := mergeSort(nums, s, m)
	n += mergeSort(nums, m+1, e)
	n += merge(nums, s, m, e)

	return n
}

func merge(nums []int, s, m, e int) int {
	i := s
	j := m+1
	n := 0

	for i <= m {
		for j <= e && nums[i] > 2*nums[j] {
			j++
		}

		n += j - (m+1)
		i++
	}


	tmp := make([]int, 0, e-s+1)
	i = s
	j = m+1

	for i <= m && j <= e {
		if nums[i] < nums[j] {
			tmp = append(tmp, nums[i])
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}

	for i <= m {
		tmp = append(tmp, nums[i])
		i++
	}

	for j <= e {
		tmp = append(tmp, nums[j])
		j++
	}

	for i, v := range tmp {
		nums[i+s] = v
	}

	// fmt.Println()
	return n
}
