package main

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	slow := head
	fast := head.Next
	for fast != nil {
		if fast.Val == val {
			slow.Next = fast.Next
		} else {
			slow = slow.Next
		}
		fast = fast.Next
	}
	if head.Val == val {
		head = head.Next
	}
	return head
}
