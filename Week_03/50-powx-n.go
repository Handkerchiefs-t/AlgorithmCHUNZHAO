package Week_03

func myPow(x float64, n int) float64 {
	if n < 0 {
		return recursion(1/x, -n)
	}

	return recursion(x, n)
}

func recursion(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	r := recursion(x, n/2)

	if n % 2 == 1 {
		return r*r*x
	}

	return r*r
}

