package main

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	node := head
	for node != nil && node.Next != nil {
		prev := dummy
		tmp := prev.Next
		for tmp != node.Next {
			if tmp.Val >= node.Next.Val {
				next := node.Next.Next
				prev.Next = node.Next
				node.Next.Next = tmp
				node.Next = next
				break
			} else {
				prev = prev.Next
				tmp = prev.Next
			}
		}
		node = node.Next
	}
	return dummy.Next
}
