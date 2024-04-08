package easy

import (
	"strconv"
)

func pad(s string, n int) string {
		for i := 0; i < n; i++ {
				s = "0" + s
		}

		return s
}

func AddBinary(a string, b string) string {
		if len(a) > len(b) {	
				b = pad(b, len(a) - len(b))
		} else {
				a = pad(a, len(b) - len(a))
		}

		carry := "0"
		result := ""
		for i := len(a) - 1; i >= 0; i-- {
				count := 0
				if a[i] == '1' {
						count++
				}
				
				if b[i] == '1' {
						count++
				}

				if carry == "1" {
						count++
				}

				if count >= 2 {
						carry = "1"
				} else {
						carry = "0"
				}

				if count % 2 == 1 {
						result = "1" + result
				} else {
						result = "0" + result
				}
		}

		if carry == "1" {
				result = "1" + result
		}

		return result
}

// This is the easy solution to the problem using strconv
func AddBinaryBonus(a string, b string) (string, error) {
		aInt, err := strconv.ParseInt(a, 2, 64)
		if err != nil {
				return "", err 
		}

		bInt, err := strconv.ParseInt(b, 2, 64)
		if err != nil {
				return "", err
		}

		sum := aInt + bInt

		binary := strconv.FormatInt(sum, 2)

		return binary, nil
}
