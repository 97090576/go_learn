package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	res := head
	for head.Next != nil && head.Next.Next != nil {
		head = head.Next.Next
		res = res.Next
	}
	if head.Next != nil {
		res = res.Next
	}
	return res
}
