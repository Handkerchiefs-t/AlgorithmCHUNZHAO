package Week_06
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	counter := [26]int{}
	tmpCounter := [26]int{}
	res := []int{}

	for _, v := range p {
		counter[v-'a']++
	}

	for i := 0; i < len(p); i++ {
		tmpCounter[s[i]-'a']++
	}

	begin := 0
	end := begin + len(p) - 1
	for {
		if counter == tmpCounter {
			res = append(res, begin)
		}

		tmpCounter[s[begin]-'a']--

		begin++
		end++

		if end == len(s) {
			break
		}

		tmpCounter[s[end]-'a']++
	}

	return res
}

