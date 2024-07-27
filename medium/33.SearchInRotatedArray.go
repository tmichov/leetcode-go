package medium

import "fmt"

func SearchInRotatedArray(nums []int, target int) int {
	lo := 0
	hi := len(nums)-1

	indx := -1

	for lo <= hi {
		mid := (lo + hi) / 2

		if nums[mid] == target {
			indx = mid
			break
		}

		fmt.Println(lo, hi, mid)
		if nums[lo] <= nums[mid] {
			if nums[lo] <= target && nums[mid] > target {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}

	}

	return indx
}

