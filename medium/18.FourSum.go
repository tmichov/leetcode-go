package medium

import (
	"fmt"
	"slices"
)

func FourSum(nums []int, target int) [][]int {
	slices.Sort(nums)

	res := make([][]int, 0)

	size := len(nums)

	for i := 0; i < size - 3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i+1; j < size - 2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			l, r := j+1, size - 1

			for l < r {

				sum := nums[i] + nums[j] + nums[l] + nums[r]

				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})

					for l < r && nums[l] == nums[l+1] {
						l++
					}

					for r > l && nums[r] == nums[r-1] {
						r--
					}

					l++
					r--
				} else if sum < target {
					l++
				} else if sum > target {
					r--
				}

			}
		}
	}

	fmt.Println(res)

	return res
}
