package main

import (
	"fmt"

	"github.com/tmichov/leetcode/easy"
)

func main() {
	node := &easy.TreeNode{
		Val: 1,
	}

	node.Right = &easy.TreeNode{Val: 2}
	node.Left = &easy.TreeNode{Val: 2}

	node.Right.Left = &easy.TreeNode{Val: 2}
	node.Right.Right = &easy.TreeNode{Val: 2}

	balanced := easy.MinimumDepthOfBinaryTree(node)

	fmt.Println("balanced: ", balanced)
}

