package easy

import "fmt"

//88. Merge sorted arrays
func MergeSortedArrays(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	for i := 0; i < m; i++ {
		if nums1[i] > nums2[0] {
			c := nums1[i]
			nums1[i] = nums2[0]

			foundAt := 0
			for j := foundAt; j < n; j++ {
				foundAt = j
				if nums2[j] > c {
					nums2 = append(nums2[1:j], append([]int{c}, nums2[j:]...)...)
					fmt.Println(nums2)
					break;
				}

				if j == len(nums2)-1 {
					nums2 = append(nums2[1:], c)
				}
			}
		}
	}

	for i2, val := range nums2 {
		nums1[m+i2] = val
	}

	fmt.Println("nums1", nums1)
}
