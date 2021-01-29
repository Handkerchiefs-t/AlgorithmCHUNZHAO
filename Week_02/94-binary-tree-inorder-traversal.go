package Week_02

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	recursion(root, &res)
	return res
}

func recursion(root *TreeNode, res *[]int) {
	if root != nil {
		recursion(root.Left, res)
		*res = append(*res, root.Val)
		recursion(root.Right, res)
	}
}
