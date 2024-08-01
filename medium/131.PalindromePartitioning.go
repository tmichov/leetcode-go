package medium

func PalindromePartition(s string) [][]string {
	res := [][]string{}

	generatePalindromePartitions(0, s, []string{}, &res)

	return res
}

func generatePalindromePartitions(index int, s string, combination []string, res *[][]string) {

	if index == len(s) {
		combCopy := append([]string{}, combination...)

		*res = append(*res, combCopy)
		return
	}

	for i := index; i < len(s); i++ {
		if isPalindrome(s[index:i+1]) {
			combination = append(combination, s[index:i+1])
			generatePalindromePartitions(i+1, s, combination, res)
			combination = combination[:len(combination)-1]
		}
	}
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false;
		}
	}

	return true
}

