package Week_06

func validPalindrome(s string) bool {
	bList := []byte(s)

	r0, l, r := isPalindrome(bList, 0, len(bList)-1)
	if r0 {
		return true
	}

	r1, _, _ := isPalindrome(bList, l+1, r)
	r2, _, _ := isPalindrome(bList, l, r-1)

	return r1 || r2
}

func isPalindrome(arr []byte, l, r int) (bool, int, int) {
	for l < r {
		if arr[l] != arr[r] {
			return false, l, r
		}

		l++
		r--
	}

	return true, l, r
}
