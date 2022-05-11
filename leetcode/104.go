package main

func maxDepth(root *TreeNode) int {
	var f func(*TreeNode, int) int
	f = func(node *TreeNode, i int) int {
		if node != nil {
			i++
			i = max(f(node.Left, i), f(node.Right, i))
		}
		return i
	}
	return f(root, 0)
}
