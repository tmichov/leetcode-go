package easy

func MajorityElements(nums []int) int {
	me, c := nums[0], 1

	for _, num := range nums[1:] {
		if num != me {
			c -= 1

			if c == 0 {
				me = num
				c = 1
			}
		} else {
			c += 1
		}
	}

	return me
}
