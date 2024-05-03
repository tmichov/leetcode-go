package easy

func ConvertSorterArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	 
	if len(nums) == 1 {
		return &TreeNode{
			Val: nums[0],
		}
	}

	root := &TreeNode{
		Val: nums[len(nums)/2],
		Left: ConvertSorterArrayToBST(nums[:len(nums)/2]),
		Right: ConvertSorterArrayToBST(nums[len(nums)/2+1:]),
	}

	return root
}

