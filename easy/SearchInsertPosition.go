package easy

func SearchIndexPosition(nums []int, target int) int {
		if nums[0] >= target {
				return 0
		}

		l := len(nums)
		if nums[l-1] <= target {
				return l
		}

		start := 0
		end := l
		for {
				check := (start+end)/2        

				if nums[check] == target {
						return check
				} else if nums[check] > target {
						end = check
				} else {
						start = check
				}

				if end - start == 1 {
						return end
				}
		}
}
