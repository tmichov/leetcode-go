package easy

import "github.com/tmichov/leetcode/util"

func IsBalanced(root *util.TreeNode) bool {
    if root == nil{
        return true
    }
    dif := maxDepth(root.Left) - maxDepth(root.Right)
    if dif <= 1 && dif >= -1{
        return IsBalanced(root.Left) && IsBalanced(root.Right)
    }
    return false
}

func maxDepth(root *util.TreeNode) int{
    if root == nil{
        return 0
    }
    return max(maxDepth(root.Left),maxDepth(root.Right))+1
}

func max(x int, y int) int{
    if x > y{
        return x
    }
    return y
}

