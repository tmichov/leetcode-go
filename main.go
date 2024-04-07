package main

import (
	"fmt"

	"github.com/tmichov/leetcode/easy"
)

func main() {
	
	nums1 := []int{0}

	nums2 := []int{1}

	easy.MergeSortedArrays(nums1, len(nums1)-len(nums2), nums2, len(nums2))

	fmt.Println()
}

