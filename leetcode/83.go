package main

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 一前一后，也就是说 slow.Next == fast 恒为 TRUE
	slow := head
	fast := head.Next
	for fast != nil {
		if fast.Val == slow.Val {
			slow.Next = fast.Next
		} else {
			slow = slow.Next
		}
		fast = fast.Next
	}
	return head
}
