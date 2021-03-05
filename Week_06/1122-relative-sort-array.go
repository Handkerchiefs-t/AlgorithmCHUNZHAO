package Week_06

func relativeSortArray(arr1 []int, arr2 []int) []int {
	maxValue := 0
	for _, v := range arr1 {
		if v > maxValue {
			maxValue = v
		}
	}

	counter := make([]int, maxValue+1)

	for _, v := range arr1 {
		counter[v]++
	}

	res := make([]int, 0, len(arr1))

	for _, v := range arr2 {
		for counter[v] > 0 {
			counter[v]--
			res = append(res, v)
		}
	}

	for v, count := range counter {
		for count > 0 {
			count--
			res = append(res, v)
		}
	}

	return res

}
