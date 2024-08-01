package medium
import "slices"

func CombinationSum2(candidates []int, target int) [][]int {
	slices.Sort(candidates)

	res := [][]int{}

	generateCombinationSums(0, 0, []int{}, candidates, target, &res)
	return res
}

func generateCombinationSums(index int, sum int, current []int, candidates []int, target int, res *[][]int) {
	if sum == target {
		curCopy := append([]int{}, current...)

		*res = append(*res, curCopy)
		return
	}

	if sum > target {
		return
	}

	for i := index; i < len(candidates); i++ {
		if i > index && candidates[i] == candidates[i-1] {
			continue;
		}

		sum += candidates[i]
		current = append(current, candidates[i])

		generateCombinationSums(i+1, sum, current, candidates, target, res)

		current = current[:len(current)-1]
		sum -= candidates[i]
	}
}


