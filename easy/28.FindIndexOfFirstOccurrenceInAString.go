package easy

// I know this couldve been done with brute force but I wanted to practice
// KMP alg

func FindIndexOfFirstOccurrenceInAString(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	if len(needle) > len(haystack) {
		return -1
	}

	lps := make([]int, len(needle))

	i := 1
	prevLSP := 0

	for i < len(needle) {
		if needle[i] == needle[prevLSP] {
			lps[i] = prevLSP + 1
			prevLSP++
			i++
		} else if prevLSP == 0 {
			lps[i] = 0
			i++
		} else {
			prevLSP = lps[prevLSP-1]
		}
	}

	i = 0
	j := 0

	for i < len(haystack) {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			if j == 0 {
				i += 1
			} else {
				j = lps[j-1]
			}
		}

		if j == len(needle) {
			return i - len(needle)
		}
	}
	return -1
}
