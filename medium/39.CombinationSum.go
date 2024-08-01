package medium

func CombinationSum(candidates []int, target int) [][]int {
	res := [][]int{}

	backtrackingSum(0, target, []int{}, &res, candidates)

	return res
}

func backtrackingSum(start, target int, combination []int, res *[][]int, candidates []int) {
	if target == 0 {
		combCopy := append([]int{}, combination...);
		*res = append(*res, combCopy)
		return
	}

	if target < 0 {
		return
	}

	for i := start; i < len(candidates); i++ {
		candidate := candidates[i]
		combination = append(combination, candidate)
		backtrackingSum(i, target-candidate, combination, res, candidates)
		combination = combination[:len(combination)-1]
	}
}


