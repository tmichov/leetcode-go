package easy

import "github.com/tmichov/leetcode/util"

type Res struct {
    Res []int
}

func postorderTraversal(root *util.TreeNode) []int {
    r := Res{}
    r.traversal(root)

    return r.Res
}

func (r *Res) traversal(node *util.TreeNode) {
    if node != nil {
        r.traversal(node.Left)
        r.traversal(node.Right)

        r.Res = append(r.Res, node.Val)
    }
}
