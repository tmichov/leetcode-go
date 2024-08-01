package medium

import (
	"fmt"
	"strconv"
)

func CountAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	return generateRLE(CountAndSay(n-1))
}

func generateRLE(str string) string {

	if len(str) == 1 {
		return "1"+str;
	}

	res := ""

	count := 1
	el := fmt.Sprintf("%c", str[0])

	for i := 1; i < len(str); i++ {
		char := fmt.Sprintf("%c", str[i])

		if el == char {
			count++
		}

		if el != char {
			res += strconv.Itoa(count) + el
			count = 1
			el = char
		}
	}

	res += strconv.Itoa(count) + el

	return string(res)
}

