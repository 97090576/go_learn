package main

func flatten(root *TreeNode) {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return
	}
	if root.Left == nil {
		flatten(root.Right)
		return
	}
	if root.Right == nil {
		flatten(root.Left)
		root.Right = root.Left
		root.Left = nil
		return
	}
	right := root.Right
	left := root.Left
	flatten(left)
	flatten(right)
	for left.Right != nil {
		left = left.Right
	}
	left.Right = right
	root.Right = root.Left
	root.Left = nil
}
