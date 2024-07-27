package medium

func ReverseInteger(x int) int {
	maxNumber := (1 << 31) - 1

	reversed := 0

	negative := 1
	if x < 0 {
		x *= -1
		negative = -1
	}

	for x > 0 {
		reversed = (reversed * 10) + x%10
		x = int(x / 10)

		if reversed > maxNumber || reversed < (maxNumber*-1) {
			return 0
		}
	}

	return reversed * negative
}
