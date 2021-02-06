package Week_03

func lemonadeChange(bills []int) bool {
	five, ten := 0, 0

	for _, value := range bills {
		if value == 5 {
			five++
		} else if value == 10 {
			if five > 0 {
				five--
				ten++
			} else {
				return false
			}
		} else {
			if ten > 0 && five > 0 {
				ten--
				five--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}

	return true
}
