package easy

func PlusOne(digits []int) []int {
	l := len(digits)

	if l == 0 {
		return []int{0}
	}

	found := false
	for i := l - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			break
		} else {
			if i == 0 {
				digits[i] = 1
				found = true
			} else {
				digits[i] = 0
			}
		}
	}

	if found {
		return append(digits, 0)
	}

	return digits
}
