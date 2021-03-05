package Week_06

func reverseOnlyLetters(s string) string {
	bList := []byte(s)
	l := 0
	r := len(bList)-1

	for l < r {
		if isAlphabet(bList[l]) == false {
			l++
			continue
		}

		if isAlphabet(bList[r]) == false {
			r--
			continue
		}

		bList[l], bList[r] = bList[r], bList[l]
		l++
		r--
	}

	return string(bList)
}

func isAlphabet(i byte) bool {
	if ('a' <= i && i <= 'z') || ('A' <= i && i <= 'Z') {
		return true
	}

	return false
}
