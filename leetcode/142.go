package main

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	low, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		low = low.Next
		if fast == low {
			break
		}
	}
	m := make(map[*ListNode]bool)
	if fast == low {
		for !m[fast] {
			m[fast] = true
			fast = fast.Next
		}
	}
	res := head
	for !m[res] && res != nil {
		res = res.Next
	}
	return res
}
