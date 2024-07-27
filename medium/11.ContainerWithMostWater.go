package medium

func ContainerWithMostWater(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		area := min(height[left], height[right]) * (right - left)

		if maxArea < area {
			maxArea = area
		}

		if height[left] > height[right] {
			right -= 1
		} else {
			left += 1
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

