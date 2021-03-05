package Week_06

func reverseWords(s string) string {
	bList := []byte(s)

	i := 0
	for j, v := range bList {
		if v != ' ' {
			i = j
			break
		}
	}

	bList = bList[i:]

	for i = len(bList)-1; i >= 0; i-- {
		if bList[i] != ' ' {
			break
		}
	}

	bList = bList[:i+1]

	for i := 0; i < len(bList)-1; {
		if bList[i] == ' ' && bList[i+1] == ' ' {
			bList = append(bList[:i], bList[i+1:]...)
			continue
		}

		i++
	}

	l := 0
	r := len(bList)-1

	for l < r {
		bList[l], bList[r] = bList[r], bList[l]
		l++
		r--
	}

	l = 0
	for i := 0; i < len(bList); i++ {
		if i == len(bList)-1 || bList[i+1] == ' ' {
			r = i

			for l < r {
				bList[l], bList[r] = bList[r], bList[l]
				l++
				r--
			}

			l = i + 2
		}
	}

	return string(bList)
}
