package Week_01

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	recursion(l1, l2, res)
	return res.Next
}

func recursion(a, b, c *ListNode) {
	if a != nil && b != nil {
		if a.Val <= b.Val {
			c.Next = a
			a = a.Next
			c = c.Next
			c.Next = nil
		} else {
			c.Next = b
			b = b.Next
			c = c.Next
			c.Next = nil
		}

		recursion(a, b, c)
	}

	if a == nil {
		c.Next = b
		return
	}

	if b == nil {
		c.Next = a
	}
}
