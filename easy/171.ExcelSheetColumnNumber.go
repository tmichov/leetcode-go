package easy

func ExcelSheetColumnNumber(columnTitle string) int {
	result := 0
	for _, char := range columnTitle {
		result = result*26 + int(char-'A'+1)
	}

	return result
}
