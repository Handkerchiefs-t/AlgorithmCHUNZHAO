package Week_06

func reverseStr(s string, k int) string {
	bList := []byte(s)

	for i := 0; i < len(bList); i+=2*k {
		l := i
		r := min(l+k-1, len(bList)-1)

		for l < r {
			bList[l], bList[r] = bList[r], bList[l]
			l++
			r--
		}
	}

	return string(bList)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
