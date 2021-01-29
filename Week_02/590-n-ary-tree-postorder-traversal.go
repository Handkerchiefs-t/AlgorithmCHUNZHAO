package Week_02

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
	res := []int{}
	recursion(root, &res)
	return res
}

func recursion(root *Node, res *[]int) {
	if root != nil {
		for _, item := range root.Children {
			recursion(item, res)
		}
		*res = append(*res, root.Val)
	}
}
