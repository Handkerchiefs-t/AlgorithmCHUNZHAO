package Week_06

func firstUniqChar(s string) int {
	counter := [26]int{}

	for _, v := range s {
		counter[v-'a']++
	}

	for i, v := range s {
		if counter[v-'a'] == 1 {
			return i
		}
	}

	return -1
}
