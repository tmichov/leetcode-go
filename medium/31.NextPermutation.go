package medium

import (
	"fmt"
)

func NextPermutation(nums []int) {
	length := len(nums)-1
	largeIndex := length
	peakIndex := length

	for ; peakIndex > 0; peakIndex-- {
		if nums[peakIndex-1] < nums[peakIndex] {
			break
		}
	}

	fmt.Println(peakIndex)

	if peakIndex != 0 {
		for ; largeIndex >= peakIndex; largeIndex-- {
			if nums[peakIndex-1] < nums[largeIndex] {
				nums[largeIndex], nums[peakIndex-1] = nums[peakIndex-1], nums[largeIndex]
				break
			}
		}
	}

	for i:=length; i>peakIndex; i, peakIndex = i-1, peakIndex+1 {
		nums[i], nums[peakIndex] = nums[peakIndex], nums[i]
	}

	fmt.Println(nums)
}
