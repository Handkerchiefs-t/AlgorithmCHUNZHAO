package Week_06

func reverseWords(s string) string {
	bList := []byte(s)

	l := 0
	for i := 0; i < len(s); i++ {
		if i == len(bList)-1 || bList[i+1] == ' ' {
			r := i

			for l < r {
				bList[l], bList[r] = bList[r], bList[l]
				l++
				r--
			}

			l = i+2
		}
	}

	return string(bList)
}
