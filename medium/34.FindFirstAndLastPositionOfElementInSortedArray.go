package medium

import "fmt"

func FindFirstAndLastPositionOfElementInSortedArray(nums []int, target int) []int {
	res := []int{-1,-1}

	if len(nums) == 0 {
		return res
	}

	res[0] = binarySearchFindMin(nums, target, true)
	res[1] = binarySearchFindMin(nums, target, false)

	return res
}

func binarySearchFindMin(nums []int, target int, findingMin bool) int {
	res := -1

	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right - left)/2

		if nums[mid] == target {
			if res == -1 {
				res = mid
			} else {
				if findingMin && mid < res {
					res = mid
				} else if !findingMin && mid > res {
					res = mid
				}
			}

			if findingMin {
				right = mid - 1
			} else {
				left = mid + 1
			}

		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}

		fmt.Println(mid)
	}


	return res
}
