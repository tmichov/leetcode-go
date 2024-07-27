package medium

import (
	"fmt"
	"slices"
)

func ThreeSum(nums []int) [][]int {
	slices.Sort(nums)

	res := make([][]int, 0)

	for i := 0; i < len(nums)-1; i++ {
		j, k := i+1, len(nums)-1;

		if j == k {
			break
		}

		if k == 0 {
			break
		}

		target := -nums[i]

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j < k {
			if nums[j] + nums[k] == target {
				res = append(res, []int{nums[i], nums[j], nums[k]})

				for j < k && nums[j] == nums[j+1] {
					j++
				}

				for j < k && nums[k] == nums[k-1] {
					k--
				}

				j++
				k--
			} else if nums[j] + nums[k] < target {
				j++
			} else {
				k--
			}
		}
	}

	fmt.Println(res)

	return nil
}
