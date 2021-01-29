package Week_02

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	return recursion(preorder, inorder)
}

// 一个非常自然的想法是，可以使用很多个参数来记录递归的状态
// 但是题解中有人用切割slice的方法省去了这些参数，真的是秀
func recursion(pre, mid []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}

	root := &TreeNode{pre[0], nil, nil}

	// 寻找 root 在 inorder 中的索引 i
	i := 0
	for ; i < len(mid); i++{
		if pre[0] == mid[i] {
			break
		}
	}

	// [root][root.Left][root.Right] <- preorder
	// [root.Left][root][root.Right] <- inorder
	root.Left = recursion(pre[1:len(mid[:i])+1], mid[:i])
	root.Right = recursion(pre[len(mid[:i])+1:], mid[i+1:])

	return root
}
