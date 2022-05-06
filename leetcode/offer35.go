package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	// 1. 每个节点拷贝一份
	for node := head; node != nil; node = node.Next.Next {
		node.Next = &Node{node.Val, node.Next, nil}
	}
	// 2. 拷贝 Random 指针
	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}
	// 3. 断开连接形成新的链表同时还原原来的链表
	headNew := head.Next
	for node := head; node != nil; node = node.Next {
		nodeNew := node.Next
		node.Next = node.Next.Next
		if nodeNew.Next != nil {
			nodeNew.Next = nodeNew.Next.Next
		}
	}
	return headNew
}
