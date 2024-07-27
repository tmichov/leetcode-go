package medium

import "fmt"

func DivideTwoIntegers(dividend, divisor int) int {
	sign := true

	if divisor == 1 {
		return dividend
	}


	if (dividend < 0 && divisor > 0) || (dividend >= 0 && divisor < 0) {
		sign = false
	}

	dividend = abs(dividend)
	divisor = abs(divisor)
	result := 0

	for dividend >= divisor {
		temp := divisor
		count := 1


		for dividend >= (temp << 1) {
			temp <<= 1
			count <<= 1

			fmt.Println("temp", temp)
			fmt.Println("count", count)
		}

		dividend -= temp
		result += count
	}

	if !sign {
		return -result
	}

	if result > 2147483647 {
		return 2147483647
	}

	return result
}

