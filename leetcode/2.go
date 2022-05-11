package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var f func(*ListNode, *ListNode) *ListNode
	cnt := 0
	f = func(node *ListNode, node2 *ListNode) *ListNode {
		if cnt == 0 {
			if node == nil {
				return node2
			}
			if node2 == nil {
				return node
			}
		} else {
			l := &ListNode{Val: cnt}
			if node == nil && node2 == nil {
				return l
			} else if node == nil {
				cnt = 0
				node = l
			} else if node2 == nil {
				cnt = 0
				node2 = l
			}
		}
		i := node.Val + node2.Val + cnt
		node.Val = i % 10
		cnt = i / 10
		node.Next = f(node.Next, node2.Next)
		return node
	}
	f(l1, l2)
	return l1
}
