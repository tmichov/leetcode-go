package medium

import (
	"fmt"
	"math"
	"slices"
)

func ThreeSumClosest(nums []int, target int) int {
	slices.Sort(nums)

	diff := math.MaxInt32
	res := 0

	for i := 0; i < len(nums)-1; i++ {
		j, k := i+1, len(nums)-1;

		if j > k {
			break
		}

		if k == -1 {
			break
		}

		fmt.Println(i,j,k)
		for j < k {
			sum := nums[i] + nums[j] + nums[k]

			if sum == target {
				return sum
			}

			if abs(target-sum) < diff {
				diff = abs(target-sum)
				res = sum
			}

			if sum > target {
				k--
			} else {
				j++
			}
		}
	}

	fmt.Println("diff", diff, "res", res)
	return res
}

func abs(i int) int {

	if i < 0 {
		return i * -1
	}

	return i
}
