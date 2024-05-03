package main

import (
	"fmt"

	"github.com/tmichov/leetcode/easy"
)

func main() {
	node := &easy.TreeNode{
		Val: 1,
	}

	node.Left = &easy.TreeNode{Val: 2}
	node.Right = &easy.TreeNode{Val: 2}

	node.Left.Left = &easy.TreeNode{Val: 2}

	node.Right.Right = &easy.TreeNode{Val: 2}

	node.Left.Left.Left = &easy.TreeNode{Val: 2}
	node.Right.Right.Right = &easy.TreeNode{Val: 2}
	balanced := easy.IsBalanced(node)

	fmt.Println("balanced: ", balanced)
}

