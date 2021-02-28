package Week_05

func isPowerOfTwo(n int) bool {
	if n <= 0 || n &(n-1) != 0 {
		return false
	}

	return true
}
