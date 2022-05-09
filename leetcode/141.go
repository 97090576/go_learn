package main

// 快慢指针，龟兔赛跑算法，跑的快的兔子能追上跑的慢的乌龟，就说明有环
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
