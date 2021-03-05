package Week_06

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func (i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	s := intervals[0][0]
	e := intervals[0][1]
	res := [][]int{}

	for i := 1; i < len(intervals); i++ {
		if e >= intervals[i][0] {
			e = max(e, intervals[i][1])
		} else {
			res = append(res, []int{s, e})
			s = intervals[i][0]
			e = intervals[i][1]
		}
	}

	res = append(res, []int{s, e})

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
