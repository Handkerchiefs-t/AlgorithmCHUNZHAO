package Week_02

// 丑数*丑数 = 丑数

func nthUglyNumber(n int) int {
	arr := []int{1}
	a, b, c := 0, 0, 0

	for i := 1; i <= n-1; i++{
		// 递推公式: min(arr[a]*2, arr[b]*3, arr[c]*5
		arr = append(arr, min(arr[a]*2, arr[b]*3, arr[c]*5))

		if arr[i] == arr[a]*2 {a++}
		if arr[i] == arr[b]*3 {b++}
		if arr[i] == arr[c]*5 {c++}
	}

	return arr[len(arr)-1]
}

func min(arr ...int) int {
	m := arr[0]

	for _, value := range arr {
		if value < m {
			m = value
		}
	}

	return m
}
