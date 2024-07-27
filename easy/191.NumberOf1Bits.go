package easy

func NumberOf1Bits(n int) int {
	result := 0
	for n > 0 {
		if (n & 1) == 1 {
			result += 1
		}

		n >>= 1
	}

	return result
}
