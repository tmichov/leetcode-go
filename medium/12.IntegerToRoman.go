package medium

func IntegerToRoman(n int) string {
	symbol := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	value := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	result := ""

	for n > 0 {
		for i := range value {
			if n >= value[i] {
				result += symbol[i]
				n -= value[i]
				break
			}
		}
	}
	return result
}
