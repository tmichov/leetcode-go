package easy

func SingleNumber(nums []int) int {
	result := 0

	for _, v := range nums {
		result ^= v
	}

	return result
}
