package medium

func LetterCombinationsOfAPhoneNumber(digits string) []string {
	if digits == "" {
		return []string{}
	}

	letters := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	combinations := []string{""}

	for _, digit := range digits {
		newCombinations := []string{}

		for _, combination := range combinations {
			for _, letter := range letters[digit-'2'] {
				newCombinations = append(newCombinations, combination+string(letter))
			}
		}

		combinations = newCombinations
	}

	return combinations
}
