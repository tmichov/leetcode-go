package easy

import "github.com/tmichov/leetcode/util"

func ConvertSorterArrayToBST(nums []int) *util.TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &util.TreeNode{
			Val: nums[0],
		}
	}

	root := &util.TreeNode{
		Val:   nums[len(nums)/2],
		Left:  ConvertSorterArrayToBST(nums[:len(nums)/2]),
		Right: ConvertSorterArrayToBST(nums[len(nums)/2+1:]),
	}

	return root
}
