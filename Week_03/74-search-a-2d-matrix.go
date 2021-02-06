package Week_03

func searchMatrix(matrix [][]int, target int) bool {
	left := 0
	right := len(matrix)*len(matrix[0])-1

	for left <= right {
		midId := left + (right-left)/2
		midValue := matrix[midId/len(matrix[0])][midId%len(matrix[0])]

		if midValue == target {
			return true
		}

		if midValue < target {
			left = midId + 1
		} else {
			right = midId - 1
		}
	}

	return false
}

