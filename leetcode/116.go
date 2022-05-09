package main

type Node116 struct {
	Val   int
	Left  *Node116
	Right *Node116
	Next  *Node116
}

// 可以使用上一层已经建立连接的指针
func connect(root *Node116) *Node116 {
	if root == nil {
		return nil
	}
	for l := root; l.Left != nil; l = l.Left {
		for n := l; n != nil; n = n.Next {
			n.Left.Next = n.Right
			if n.Next != nil {
				n.Right.Next = n.Next.Left
			}
		}
	}
	return root
}
