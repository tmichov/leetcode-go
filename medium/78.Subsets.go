package medium

func Subsets(nums []int) [][]int {

	res := [][]int{}

	generateSubsets(nums, 0, []int{}, &res)

	return res
}

func generateSubsets(set []int, index int, current []int, res *[][]int) {
	if index == len(set) {
		curCopy := append([]int{}, current...)

		*res = append(*res, curCopy)
		return
	}

	generateSubsets(set, index+1, current, res)

	current = append(current, set[index])

	generateSubsets(set, index+1, current, res)

	current = current[:len(current)-1]
}
