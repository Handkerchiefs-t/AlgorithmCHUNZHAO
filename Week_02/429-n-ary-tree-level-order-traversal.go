package Week_02

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	tmp := []int{}
	queue := []*Node{root, nil}

	for {
		tmpNode := queue[0]
		queue = queue[1:]

		if tmpNode == nil {
			res = append(res, tmp)

			if len(queue) == 0 {break}
			tmp = []int{}
			queue = append(queue, nil)
		} else {
			tmp = append(tmp, tmpNode.Val)
			for _, i := range tmpNode.Children {
				queue = append(queue, i)
			}
		}
	}

	return res
}