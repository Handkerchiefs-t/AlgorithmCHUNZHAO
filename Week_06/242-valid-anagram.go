package Week_06

func isAnagram(s string, t string) bool {
	counter1 := [26]int{}
	counter2 := [26]int{}

	for _, v := range s {
		counter1[v-'a']++
	}

	for _, v := range t {
		counter2[v-'a']++
	}

	for i, v := range counter1 {
		if counter2[i] != v {
			return false
		}
	}

	return true
}
