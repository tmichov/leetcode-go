package medium

import (
	"strings"
)

func ZigZagConversion(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	matrix := make([]string, numRows)

	row := 0
	up := true

	for i := 0; i < len(s); i++ {
		matrix[row] += string(s[i])

		if row == 0 || row == numRows-1 {
			up = !up
		}

		if up {
			row -= 1
		} else {
			row += 1
		}
	}

	return strings.Join(matrix, "")
}
