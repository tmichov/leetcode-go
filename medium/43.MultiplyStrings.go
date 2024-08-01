package medium

func MultiplyStrings(num1, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	res := make([]byte, len(num1) + len(num2))

	for i := len(num1)-1; i >= 0; i-- {
		for j := len(num2)-1; j >= 0; j-- {
			v1 := num1[i]-'0'
			v2 := num2[j]-'0'

			res[i+j+1] += (v1 * v2);

			if res[i+j+1] >= 10 {
				res[i+j] += (res[i+j+1]/10)
				res[i+j+1] %= 10
			}
		}
	}

	if res[0] == 0 {
		res = res[1:]
	}

	for i := range res {
		res[i] = res[i]+'0'
	}

	return string(res)
}


