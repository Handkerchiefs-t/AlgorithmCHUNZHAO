package Week_01

func groupAnagrams(strs []string) [][]string {
	m := map[[26]int][]string{}
	for _, s := range strs {
		tmp := [26]int{}
		for _, a := range s {
			tmp[a-'a']++
		}
		m[tmp] = append(m[tmp], s)
	}

	res := [][]string{}
	for _, i := range m {
		res = append(res, i)
	}
	return res
}
