package easy

func excelSheetColumnTitle(n int) string {
	if n == 0 {
		return ""
	}

	return excelSheetColumnTitle((n-1)/26) + string(('A' + byte((n-1)%26)))
}
