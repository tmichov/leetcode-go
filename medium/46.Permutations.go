package medium

func Permute(nums []int) [][]int {
	res := [][]int{}

	generatePermutations(0, nums, &res)

	return res
}

func generatePermutations(index int, set []int, res *[][]int) {
	if index == len(set) {
		curCopy := append([]int{}, set...)

		*res = append(*res, curCopy)
		return
	}

	for i := index; i < len(set); i++ {
		set[index], set[i] = set[i], set[index]

		generatePermutations(index + 1, set, res)

		set[index], set[i] = set[i], set[index]
	}
}
